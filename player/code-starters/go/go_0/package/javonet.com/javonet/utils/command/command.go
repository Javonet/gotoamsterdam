package command

import (
	"fmt"
)

type Command struct {
	TargetRuntime byte
	CommandType   byte
	Payload       []interface{}
}

func (cmd *Command) String() string {
	return fmt.Sprintf("Target runtime: %d. . Command type: %d. Payload: %v", cmd.TargetRuntime, cmd.CommandType, cmd.Payload)
}

//NewCommand: creates new command.
//Args:
//tr: target runtime
//ct: command type
//p: payload
func NewCommand(tr byte, ct byte, p *[]interface{}) *Command {
	if p == nil {
		return &Command{
			TargetRuntime: tr,
			CommandType:   ct,
			Payload:       nil}
	} else {
		return &Command{
			TargetRuntime: tr,
			CommandType:   ct,
			Payload:       *p}
	}
}

func (cmd *Command) AddArgToPayload(arg *[]interface{}) *Command {
	return &Command{cmd.TargetRuntime, cmd.CommandType, append(cmd.Payload, *arg...)}
}

func (cmd *Command) PrependArgToPayload(arg *[]interface{}) *Command {
	if (arg  == nil) {
		return &Command{cmd.TargetRuntime, cmd.CommandType, cmd.Payload}
	} else 
	{
		return &Command{cmd.TargetRuntime, cmd.CommandType, append(*arg, cmd.Payload...)}	
	}
}

func (cmd *Command) DropFirstPayloadArgument() *Command {
	if len(cmd.Payload) > 0 {
		return &Command{cmd.TargetRuntime, cmd.CommandType, cmd.Payload[1:]}
	} else {
		return &Command{cmd.TargetRuntime, cmd.CommandType, cmd.Payload}
	}
}
