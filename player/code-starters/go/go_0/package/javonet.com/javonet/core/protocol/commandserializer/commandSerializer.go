package commandserializer

import (
	"fmt"

	"javonet.com/javonet/core/protocol/typeserializer"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/runtimename"
)

type Stack []command.Command

type CommandSerializer struct {
	buffer []byte
}

func NewCommandSerializer() *CommandSerializer {
	return &CommandSerializer{buffer: nil}
}

// Serialize: transform command into bytearray
func (serializer *CommandSerializer) Serialize(rootCommand *command.Command, connType byte, connData connectiondata.TcpConnectionData, runtimeVersion byte) (encodedCmd []byte, err error) {
	if rootCommand == nil {
		return nil, fmt.Errorf("command is nil")
	}
	serializer.buffer = []byte{}
	stack := new(Stack)
	stack.push(rootCommand)
	serializer.insertIntoBuffer([]byte{rootCommand.TargetRuntime, runtimeVersion, connType})
	serializer.insertIntoBuffer(connData.GetAddressBytes())
	serializer.insertIntoBuffer(connData.GetPortBytes())
	serializer.insertIntoBuffer([]byte{runtimename.Go})
	serializer.insertIntoBuffer([]byte{rootCommand.CommandType})
	encodedCmd = serializer.encodeRecursively(stack)
	return encodedCmd, nil
}

func (serializer *CommandSerializer) encodeRecursively(stack *Stack) []byte {
	if stack.isEmpty() {
		return serializer.buffer
	}
	currentCmd, _ := stack.pop()
	stack.push(currentCmd.DropFirstPayloadArgument())
	if len(currentCmd.Payload) > 0 {
		switch element := currentCmd.Payload[0].(type) {
		case *command.Command:
			innerCmd := element
			encodedCmd, _ := typeserializer.Serialize_command(innerCmd)
			serializer.insertIntoBuffer(encodedCmd)
			stack.push(innerCmd)
		default:
			encodedPrimitive, _ := typeserializer.Serialize_primitive(currentCmd.Payload[0])
			serializer.insertIntoBuffer(encodedPrimitive)
		}
	} else {
		stack.pop()
	}
	return serializer.encodeRecursively(stack)
}

func (serializer *CommandSerializer) insertIntoBuffer(arg []byte) {
	serializer.buffer = append(serializer.buffer, arg...)
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(cmd *command.Command) {
	*s = append(*s, *cmd) //Append new value to the end of the stack
}

func (s *Stack) pop() (command.Command, bool) {
	index := len(*s) - 1   //Get the index of last element
	element := (*s)[index] //Take the last element
	*s = (*s)[:index]      //Remove it from slice
	return element, true
}
