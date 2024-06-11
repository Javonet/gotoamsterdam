package handler

import (
	"fmt"

	"javonet.com/javonet/utils/command"
)

type InvokeStaticMethodHandler struct {
	AbstractCommandHandler
}

func (h *InvokeStaticMethodHandler) process(cmd *command.Command) (interface{}, error) {
	if len(cmd.Payload) < 1 {
		return false, fmt.Errorf("invokeStatciMethod payload parameters mismatch")
	}
	var err error
	var response []interface{}

	f, err = plugin_.Lookup(fmt.Sprintf("%s", cmd.Payload[0]))
	if err != nil {
		return nil, err
	}
	if len(cmd.Payload) == 1 {
		functionToInvoke := f.(func())
		functionToInvoke()
		response = nil
	} else {
		functionToInvoke := f.(func([]interface{}) []interface{})
		response = functionToInvoke([]interface{}{cmd.Payload[1:]})
	}
	return response, nil
}
