package handler

import (
	"plugin"

	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
)

type Handler struct {
	HandlerDict map[byte]*AbstractCommandHandler
}

var plugin_ *plugin.Plugin
var f plugin.Symbol
var v plugin.Symbol

func NewHandler() *Handler {
	dict := make(map[byte]*AbstractCommandHandler)
	dict[commandtype.LoadLibrary] = &AbstractCommandHandler{CmdHandler: &LoadLibraryHandler{}}
	dict[commandtype.InvokeStaticMethod] = &AbstractCommandHandler{CmdHandler: &InvokeStaticMethodHandler{}}
	dict[commandtype.GetType] = &AbstractCommandHandler{CmdHandler: &GetTypeHandler{}}
	return &Handler{HandlerDict: dict}
}

func (h *Handler) HandleCommand(cmd *command.Command) (*command.Command, error) {
	response, err := h.HandlerDict[cmd.CommandType].HandleCmd(cmd)
	if err != nil {
		return command.NewCommand(cmd.TargetRuntime, commandtype.Exception, &[]interface{}{}), err
	}
	switch response := response.(type) {
	case bool, byte, int8, int32, int64, string:
		return command.NewCommand(cmd.TargetRuntime, commandtype.Value, &[]interface{}{response}), nil
	default:
		rc := NewReferenceCache()
		uuidStr := rc.CacheReference(response)
		return command.NewCommand(cmd.TargetRuntime, commandtype.Value, &[]interface{}{uuidStr}), nil
	}
}
