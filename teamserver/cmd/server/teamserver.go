package server

import (
	"Havoc/pkg/api"
	"Havoc/pkg/colors"
	"Havoc/pkg/logger"
	"Havoc/pkg/plugin"
	"Havoc/pkg/profile"
	"os"
)

var (
	Version  = "1.0"
	CodeName = "Killer Queen"
)

func NewTeamserver() *Teamserver {
	var (
		server = new(Teamserver)
		err    error
	)

	if err = server.createConfigPath(); err != nil {
		logger.Error("couldn't create config folder: %v", err.Error())
		return nil
	}

	return server
}

func (t *Teamserver) SetServerFlags(flags TeamserverFlags) {
	t.Flags = flags
}

func (t *Teamserver) Start() {
	var (
		ServerFinished chan bool
		err            error
		server         profile.Server
	)

	if t.Server, err = api.NewServerApi(t); err != nil {
		logger.Error("failed to start api server: " + err.Error())
		return
	}

	// generate a new plugin system instance
	t.plugins = plugin.NewPluginSystem(t)

	server, err = t.profile.Server()
	if err != nil {
		logger.Error("failed to parse profile server block: " + err.Error())
		return
	}

	// load all plugins that has been specified in the folder
	logger.Info("loading plugins [%v]:", len(server.Plugins))
	for i := range server.Plugins {
		var ext *plugin.Plugin

		if ext, err = t.plugins.RegisterPlugin(server.Plugins[i]); err != nil {
			logger.Error("failed to load plugin: %v", err)
		}

		if ext != nil {
			logger.Info(" %v plugin loaded => \"%v\"", colors.BoldBlue("*"), colors.BoldBlue(ext.Name))
		}
	}

	// start the api server
	go t.Server.Start(server.Host, server.Port, t.ConfigPath(), &ServerFinished)

	logger.Info("starting server on %v", colors.BlueUnderline("https://"+server.Host+":"+server.Port))

	// This should hold the Teamserver as long as the WebSocket Server is running
	logger.Debug("wait til the server shutdown")
	<-ServerFinished

	logger.Warn("server finished (?)")
	os.Exit(0)
}

// Version
// get the current server version
func (*Teamserver) Version() map[string]string {
	return map[string]string{
		"version":  Version,
		"codename": CodeName,
	}
}

func (t *Teamserver) createConfigPath() error {
	var err error

	if t.configPath, err = os.UserHomeDir(); err != nil {
		return err
	}

	t.configPath += "/.havoc/server"

	if err = os.MkdirAll(t.configPath, os.ModePerm); err != nil {
		return err
	}

	// create listeners directory
	if err = os.MkdirAll(t.configPath+"/listeners", os.ModePerm); err != nil {
		return err
	}

	// create agents directory
	if err = os.MkdirAll(t.configPath+"/agents", os.ModePerm); err != nil {
		return err
	}

	return nil
}

func (t *Teamserver) ConfigPath() string {
	return t.configPath
}

func (t *Teamserver) Profile(path string) error {
	var err error

	t.profile = profile.NewProfile()

	logger.LoggerInstance.STDERR = os.Stderr

	err = t.profile.Parse(path)
	if err != nil {
		logger.SetStdOut(os.Stderr)
		logger.Error("Profile error: %v", colors.Red(err))
		return err
	}

	return nil
}
