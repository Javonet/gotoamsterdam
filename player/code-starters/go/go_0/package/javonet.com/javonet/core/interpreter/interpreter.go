package interpreter

import (
	"javonet.com/javonet/core/handler"
	"javonet.com/javonet/core/protocol/commanddeserializer"
	"javonet.com/javonet/core/protocol/commandserializer"
	"javonet.com/javonet/core/transmitter"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/connectiondata"
)

type Interpreter struct {
}

func NewInterpreter() *Interpreter {

	return &Interpreter{}
}

func (interpreter *Interpreter) Execute(cmd *command.Command, connType byte, connData connectiondata.TcpConnectionData) (*command.Command, error) {
	serializer := &commandserializer.CommandSerializer{}
	encodedCmd, err := serializer.Serialize(cmd, connType, connData, 0)
	if err != nil {
		emptyCommand := command.NewCommand(cmd.TargetRuntime, cmd.TargetRuntime, nil)
		return emptyCommand, err
	}
	responseByteArray, err := transmitter.SendCommand(encodedCmd)
	if err != nil {
		emptyCommand := command.NewCommand(cmd.TargetRuntime, cmd.TargetRuntime, nil)
		return emptyCommand, err
	}
	deserializer := commanddeserializer.NewCommandDeserializer(responseByteArray)
	return deserializer.Deserialize()
}

func (interpreter *Interpreter) Process(messageByteArray []byte) ([]byte, error) {
	deserializer := commanddeserializer.NewCommandDeserializer(messageByteArray)
	receivedCommand, err := deserializer.Deserialize()
	if err != nil {
		return nil, err
	}
	handler := handler.NewHandler()
	responseCommand, err := handler.HandleCommand(receivedCommand)
	if err != nil {
		return nil, err
	}
	serializer := commandserializer.NewCommandSerializer()
	encodedCmd, err := serializer.Serialize(responseCommand, messageByteArray[2], connectiondata.TcpConnectionData{}, 0)
	if err != nil {
		return nil, err
	}
	return encodedCmd, nil
}
