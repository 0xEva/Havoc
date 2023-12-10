#ifndef HAVOCCLIENT_EVENTS_H
#define HAVOCCLIENT_EVENTS_H

#include <string>

namespace Event {

    /* user events */
    namespace user {
        static std::string login   = "user::login";
        static std::string logout  = "user::logout";
        static std::string message = "user::message";
    }

    namespace agent {
        static std::string add        = "agent::register";
        static std::string initialize = "agent::initialize";
        static std::string callback   = "agent::callback";
        static std::string console    = "agent::console";
        static std::string input      = "agent::input";
        static std::string status     = "agent::status";
        static std::string remove     = "agent::remove";
    }

    namespace listener {
        static std::string add    = "listener::register";
        static std::string start  = "listener::start";
        static std::string edit   = "listener::edit";
        static std::string stop   = "listener::stop";
        static std::string status = "listener::status";
        static std::string log    = "listener::log";
    }

    namespace payload {

    }
};


#endif //HAVOCCLIENT_EVENTS_H
