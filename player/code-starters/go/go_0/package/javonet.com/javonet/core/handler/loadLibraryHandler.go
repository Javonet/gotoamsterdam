package handler

import (
	"fmt"
	"plugin"

	"javonet.com/javonet/utils/command"
)

type LoadLibraryHandler struct {
	AbstractCommandHandler
}

func (h *LoadLibraryHandler) process(cmd *command.Command) (interface{}, error) {
	if len(cmd.Payload) != 1 {
		return false, fmt.Errorf("loadLibrary payload parameters mismatch")
	}
	var err error
	plugin_, err = plugin.Open(fmt.Sprintf("%v", cmd.Payload[0]))

	if plugin_ == nil {
		return false, fmt.Errorf("error while loading plugin")
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
