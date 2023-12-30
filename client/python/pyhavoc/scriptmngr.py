from _pyhavoc import core

import sys
import os
import importlib.util

class HcPyScriptMngrStdOutErrHandler:
    def __init__( self ):
        return

    def write( self, data ):
        core.HcScriptManagerConsoleStdOut( data )

##
## Callback function once the user decides
## to load a script into the client
##
@core.HcScriptManagerLoadScriptCallback
def _HcPyScriptLoad(
    _script_path: str
) -> None:

    script_name = os.path.splitext( os.path.basename( _script_path ) )[ 0 ]
    module_name = "HcScript." + script_name
    script_spec = importlib.util.spec_from_file_location( module_name, _script_path )
    module      = importlib.util.module_from_spec( script_spec )

    script_spec.loader.exec_module( module )

    return

##
## redirect StdOut and StdErr
##
## TODO: make own class or add param for
##       StdErr to pretty print it as red
##
# sys.stdout = HcPyScriptMngrStdOutErrHandler()
# sys.stderr = HcPyScriptMngrStdOutErrHandler()