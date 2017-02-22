/*
 * Copyright (C) 2017 Meng Shi
 */

package modules

import (
      "unsafe"
    . "worker/types"
)

const (
    STDOUT_MODULE = 0x0004
)

var stdoutModule = String{ len("stdout_module"), "stdout_module" }
var outputStdoutContext = &AbstractContext{
    	stdoutModule,
    	nil,
    	nil,
}

var	stdout = String{ len("stdout"), "stdout" }
var outputStdoutCommands = []Command{

   	{ stdout,
      MAIN_CONF|CONF_1MORE,
      stdoutBlock,
      0,
      0,
      nil },

   	NilCommand,
}

func stdoutBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
   	for m := 0; Modules[m] != nil; m++ {
				    module := Modules[m]
		      if module.Type != STDOUT_MODULE {
			         continue
		      }

		      context := (*AbstractContext)(unsafe.Pointer(module.Context))
		      if context == nil {
			         continue
		      }

		      if handle := context.Create; handle != nil {
			         this := handle(cycle)

			         if cycle.SetContext(module.Index, &this) == Error {
				            return "0"
			         }
		      }
	   }

	   if configure.SetModuleType(STDOUT_MODULE) == Error {
				    return "0"
	   }

   	if configure.Parse(cycle) == Error {
      		return "0"
	   }

	   for m := 0; Modules[m] != nil; m++ {
		      module := Modules[m]
		      if module.Type != STDOUT_MODULE {
			         continue
		      }

		      this := (*AbstractContext)(unsafe.Pointer(module.Context))
		      if this == nil {
				        continue
				    }

		      context := cycle.GetContext(module.Index)
		      if context == nil {
			         continue
		      }

  		    if init := this.Init; init != nil {
    		 	    if init(cycle, context) == "-1" {
 				           return "0"
 			        }
 		     }
	   }

   	return "0"
}

var outputStdoutModule = Module{
   	MODULE_V1,
   	CONTEXT_V1,
	   unsafe.Pointer(outputStdoutContext),
	   outputStdoutCommands,
	   OUTPUT_MODULE,
	   nil,
	   nil,
}

func init() {
   	Modules = append(Modules, &outputStdoutModule)
}