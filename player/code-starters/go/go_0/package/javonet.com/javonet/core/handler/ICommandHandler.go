package handler

import "javonet.com/javonet/utils/command"

type ICommandHandler interface {
	process(cmd *command.Command) (interface{}, error)
}
