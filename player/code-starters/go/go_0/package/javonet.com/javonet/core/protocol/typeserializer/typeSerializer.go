package typeserializer

import (
	"fmt"
	"math"

	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/stringencodingmode"
	"javonet.com/javonet/utils/types"
)

func Serialize_primitive(element interface{}) ([]byte, error) {
	switch element := element.(type) {
	case string:
		return serialize_string(element), nil
	case int32:
		return serialize_int(element), nil
	case bool:
		return serialize_bool(element), nil
	case float32:
		return serialize_float(element), nil
	case byte:
		return serialize_byte(element), nil
	case int8:
		return serialize_char(element), nil
	case int64:
		return serialize_longlong(element), nil
	case float64:
		return serialize_double(element), nil
	case uint64:
		return serialize_ulonglong(element), nil
	case uint32:
		return serialize_uint(element), nil
	case int:
		//if element is withing int32 range use serialize_int, otherwise serialize_longlong
		if element > math.MaxInt32 || element < math.MinInt32 {
			return serialize_longlong(int64(element)), nil
		}
		return serialize_int(int32(element)), nil
	default:
		return nil, fmt.Errorf("Golang: error in typeSerializer. Type not supported")
	}
}

func Serialize_command(cmd *command.Command) ([]byte, error) {
	sizeArray := Serialize_size(int32(len(cmd.Payload)))
	return []byte{
		types.JavonetCommand,
		sizeArray[0],
		sizeArray[1],
		sizeArray[2],
		sizeArray[3],
		cmd.TargetRuntime,
		cmd.CommandType,
	}, nil
}

func serialize_string(s string) []byte {
	bytes := []byte(s)
	response := make([]byte, 2)
	response[0] = types.JavonetString
	response[1] = stringencodingmode.UTF8
	response = append(response, Serialize_size(int32(len(bytes)))...)
	response = append(response, bytes...)
	return response
}

func serialize_int(v int32) []byte {
	return []byte{
		types.JavonetInteger,
		types.JavonetIntegerSize,
		byte(v>>0) & 0xFF,
		byte(v>>8) & 0xFF,
		byte(v>>16) & 0xFF,
		byte(v>>24) & 0xFF,
	}
}

func serialize_bool(v bool) []byte {
	var val byte
	if v {
		val = 1
	} else {
		val = 0
	}
	return []byte{
		types.JavonetBoolean,
		types.JavonetBooleanSize,
		val,
	}
}

func serialize_float(v float32) []byte {
	value := math.Float32bits(v)
	return []byte{
		types.JavonetFloat,
		types.JavonetFloatSize,
		byte(value>>0) & 0xFF,
		byte(value>>8) & 0xFF,
		byte(value>>16) & 0xFF,
		byte(value>>24) & 0xFF,
	}
}

func serialize_byte(v byte) []byte {
	return []byte{
		types.JavonetByte,
		types.JavonetByteSize,
		v,
	}
}

func serialize_char(v int8) []byte {
	return []byte{
		types.JavonetChar,
		types.JavonetCharSize,
		byte(v),
	}
}

func serialize_longlong(v int64) []byte {
	return []byte{
		types.JavonetLongLong,
		types.JavonetLongLongSize,
		byte(v>>0) & 0xFF,
		byte(v>>8) & 0xFF,
		byte(v>>16) & 0xFF,
		byte(v>>24) & 0xFF,
		byte(v>>32) & 0xFF,
		byte(v>>40) & 0xFF,
		byte(v>>48) & 0xFF,
		byte(v>>56) & 0xFF,
	}
}

func serialize_double(v float64) []byte {
	value := math.Float64bits(v)
	return []byte{
		types.JavonetDouble,
		types.JavonetDoubleSize,
		byte(value>>0) & 0xFF,
		byte(value>>8) & 0xFF,
		byte(value>>16) & 0xFF,
		byte(value>>24) & 0xFF,
		byte(value>>32) & 0xFF,
		byte(value>>40) & 0xFF,
		byte(value>>48) & 0xFF,
		byte(value>>56) & 0xFF,
	}
}

func serialize_ulonglong(v uint64) []byte {
	return []byte{
		types.JavonetUnsignedLongLong,
		types.JavonetUnsignedLongLongSize,
		byte(v >> 0) & 0xFF,
		byte(v >> 8) & 0xFF,
		byte(v >> 16) & 0xFF,
		byte(v >> 24) & 0xFF,
		byte(v >> 32) & 0xFF,
		byte(v >> 40) & 0xFF,
		byte(v >> 48) & 0xFF,
		byte(v >> 56) & 0xFF,
	}
}

func serialize_uint(v uint32) []byte {
	return []byte{
		types.JavonetUnsignedInteger,
		types.JavonetUnsignedIntegerSize,
		byte(v >> 0) & 0xFF,
		byte(v >> 8) & 0xFF,
		byte(v >> 16) & 0xFF,
		byte(v >> 24) & 0xFF,
	}
}

func Serialize_size(v int32) []byte {
	return []byte{
		byte(v>>0) & 0xFF,
		byte(v>>8) & 0xFF,
		byte(v>>16) & 0xFF,
		byte(v>>24) & 0xFF,
	}
}
