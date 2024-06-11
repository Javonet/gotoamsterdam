package handler

import "javonet.com/javonet/utils/command"

type AbstractCommandHandler struct {
	CmdHandler           ICommandHandler
	requiredParamsNumber byte
	HandlerDict          map[byte]*AbstractCommandHandler
}

func (a *AbstractCommandHandler) HandleCmd(cmd *command.Command) (interface{}, error) {
	a.HandlerDict = NewHandler().HandlerDict
	a.iterate(cmd)
	return a.CmdHandler.process(cmd)
}

func (a *AbstractCommandHandler) iterate(cmd *command.Command) error {
	for i, element := range cmd.Payload {
		switch element := element.(type) {
		case *command.Command:
			responseCmd, err := a.HandlerDict[cmd.CommandType].HandleCmd(element)
			if err != nil {
				return err
			}
			cmd.Payload[i] = &[]interface{}{responseCmd}
		default:
			return nil
		}
	}
	return nil
}
