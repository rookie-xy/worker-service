/*
 * Copyright (C) 2017 Meng Shi
 */

package types

import "unsafe"

type AbstractOutput struct {
    *AbstractCycle
    *AbstractFile
     output  Output
}

type Output interface {}

func NewOutput() *AbstractOutput {
    return &AbstractOutput{}
}

var output = String{ len("output"), "output" }
var outputContext = &AbstractContext{
    output,
    nil,
    nil,
}

var outputs = String{ len("outputs"), "outputs" }
var outputCommands = []Command{

    { outputs,
      MAIN_CONF|CONF_BLOCK,
      outputsBlock,
      0,
      0,
      nil },

    NilCommand,
}

func outputsBlock(configure *AbstractConfigure, command *Command, cycle *AbstractCycle) string {
	for m := 0; Modules[m] != nil; m++ {
		module := Modules[m]
		if module.Type != OUTPUT_MODULE {
			continue
		}

		context := (*AbstractContext)(unsafe.Pointer(module.Context))
		if context == nil {
			continue
		}

		if this := context.Create(cycle); this != nil {
			if *(*string)(unsafe.Pointer(uintptr(this))) == "-1" {
				return "0";
			}

			if cycle.SetContext(module.Index, &this) == Error {
				return "0"
			}
		}
	}

	if configure.Parse() == Error {
		return "0"
	}

	for m := 0; Modules[m] != nil; m++ {
		module := Modules[m]
		if module.Type != OUTPUT_MODULE {
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

	return ""
}

var outputModule = Module{
    MODULE_V1,
    CONTEXT_V1,
    unsafe.Pointer(outputContext),
    outputCommands,
    CONFIG_MODULE,
    nil,
    nil,
}

func init() {
    Modules = append(Modules, &outputModule)
}