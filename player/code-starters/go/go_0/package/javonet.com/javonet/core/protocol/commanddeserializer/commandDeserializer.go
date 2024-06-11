package commanddeserializer

import (
	"fmt"

	"javonet.com/javonet/core/protocol/typedeserializer"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/types"
)

type CommandDeserializer struct {
	buffer    []byte
	bufferLen int
	position  int
	command   *command.Command
}

func NewCommandDeserializer(messageByteArray []byte) *CommandDeserializer {
	return &CommandDeserializer{buffer: messageByteArray,
		bufferLen: len(messageByteArray),
		position:  11,
		command:   command.NewCommand(messageByteArray[0], messageByteArray[10], nil),
	}
}

func (cmdDsr *CommandDeserializer) Deserialize() (*command.Command, error) {
	for ok := true; ok; ok = !cmdDsr.isAtEnd() {
		arg, err := cmdDsr.readObject(cmdDsr.buffer[cmdDsr.position])
		if err != nil {
			return nil, err
		} else {
			cmdDsr.command = cmdDsr.command.AddArgToPayload(&[]interface{}{arg})
		}
	}
	return cmdDsr.command, nil
}

func (cmdDsr *CommandDeserializer) isAtEnd() bool {
	return cmdDsr.position == cmdDsr.bufferLen
}

func (cmdDsr *CommandDeserializer) readObject(typeNum byte) (interface{}, error) {
	switch typeNum {
	case types.JavonetCommand:
		return cmdDsr.readCommand()
	case types.JavonetString:
		return cmdDsr.readString()
	case types.JavonetInteger:
		return cmdDsr.readInt()
	case types.JavonetBoolean:
		return cmdDsr.readBool()
	case types.JavonetFloat:
		return cmdDsr.readFloat()
	case types.JavonetByte:
		return cmdDsr.readByte()
	case types.JavonetChar:
		return cmdDsr.readChar()
	case types.JavonetLongLong:
		return cmdDsr.readLong()
	case types.JavonetDouble:
		return cmdDsr.readDouble()
	case types.JavonetUnsignedLongLong:
		return cmdDsr.readULong()
	case types.JavonetUnsignedInteger:
		return cmdDsr.readUInt()
	default:
		return nil, fmt.Errorf("error in commandDeserializer readObject. Type not supported")
	}
}

func (cmdDsr *CommandDeserializer) readCommand() (*command.Command, error) {
	p := cmdDsr.position
	numberOfElementsInPayload, err := typedeserializer.Deserialize_int(cmdDsr.buffer[p+1 : p+5])
	if err != nil {
		return nil, err
	}
	runtime := cmdDsr.buffer[p+5]
	command_type := cmdDsr.buffer[p+6]
	cmdDsr.position += 7
	command := command.NewCommand(runtime, command_type, &[]interface{}{})
	return cmdDsr.readCommandRecursively(numberOfElementsInPayload, command)
}

func (cmdDsr *CommandDeserializer) readCommandRecursively(numberOfElementsInPayloadLeft int32, cmd *command.Command) (*command.Command, error) {
	if numberOfElementsInPayloadLeft == 0 {
		return cmd, nil
	} else {
		p := cmdDsr.position
		arg, err := cmdDsr.readObject(cmdDsr.buffer[p])
		if err != nil {
			return nil, err
		} else {
			cmd = cmd.AddArgToPayload(&[]interface{}{arg})
		}
		return cmdDsr.readCommandRecursively(numberOfElementsInPayloadLeft-1, cmd)
	}
}

func (cmdDsr *CommandDeserializer) readString() (string, error) {
	p := cmdDsr.position
	stringEncodingMode := cmdDsr.buffer[p+1]
	size, err := typedeserializer.Deserialize_int(cmdDsr.buffer[p+2 : p+6])
	if err != nil {
		return "", err
	}
	cmdDsr.position += 6
	p = cmdDsr.position
	cmdDsr.position += int(size)
	return typedeserializer.Deserialize_string(stringEncodingMode, cmdDsr.buffer[p:p+int(size)])
}

func (cmdDsr *CommandDeserializer) readInt() (int32, error) {
	size  := 4
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_int(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readBool() (bool, error) {
	size  := 1
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_bool(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readFloat() (float32, error) {
	size  := 4
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_float(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readLong() (int64, error) {
	size  := 8
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_longlong(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readByte() (byte, error) {
	size  := 1
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_byte(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readChar() (int8, error) {
	size  := 1
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_char(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readDouble() (float64, error) {
	size  := 8
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_double(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readUInt() (uint32, error) {
	size  := 4
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_uint(cmdDsr.buffer[p : p+size])
}

func (cmdDsr *CommandDeserializer) readULong() (uint64, error) {
	size  := 8
	cmdDsr.position += 2
	p := cmdDsr.position
	cmdDsr.position += size
	return typedeserializer.Deserialize_ulonglong(cmdDsr.buffer[p : p+size])
}
