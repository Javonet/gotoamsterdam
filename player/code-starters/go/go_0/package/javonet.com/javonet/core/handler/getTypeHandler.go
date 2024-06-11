package handler

import (
	"fmt"
	"plugin"

	"javonet.com/javonet/utils/command"
)

type GetTypeHandler struct {
	AbstractCommandHandler
}

func (h *GetTypeHandler) process(cmd *command.Command) (interface{}, error) {
	if len(cmd.Payload) != 1 {
		return nil, fmt.Errorf("getType payload parameters mismatch")
	}
	var err error
	plugin_, err = plugin.Open(fmt.Sprintf("%s", cmd.Payload[0]))
	if plugin_ == nil {
		return nil, fmt.Errorf("error while loading plugin")
	}
	if err != nil {
		return nil, err
	}
	return true, nil
}
