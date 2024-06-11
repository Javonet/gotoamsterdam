package typedeserializer

import (
	"encoding/binary"
	"fmt"
	"math"

	"javonet.com/javonet/utils/stringencodingmode"
	"javonet.com/javonet/utils/types"
)

func Deserialize_bool(encoded_bool []byte) (bool, error) {
	if len(encoded_bool) != types.JavonetBooleanSize {
		return false, fmt.Errorf("error in Deserialize_bool. Wrong number of elements in byte array")
	} else if encoded_bool[0] == 1 {
		return true, nil
	} else if encoded_bool[0] == 0 {
		return false, nil
	} else {
		return false, fmt.Errorf("error in Deserialize_bool")
	}
}

func Deserialize_string(stringEncodingMode byte, encoded_string []byte) (string, error) {
		switch stringEncodingMode {
		case stringencodingmode.ASCII:
			return string(encoded_string), nil
		case stringencodingmode.UTF8:
			return string(encoded_string), nil
		case stringencodingmode.UTF16:
			return "", fmt.Errorf("stringEncodingMode.UTF16 not implemented")
		case stringencodingmode.UTF32:
			return "", fmt.Errorf("stringEncodingMode.UTF32 not implemented")
		default:
			return "", fmt.Errorf("Deserialize_string argument out of range")
		}
}

func Deserialize_int(encoded_int32 []byte) (int32, error) {
	if len(encoded_int32) != types.JavonetIntegerSize {
		return 0, fmt.Errorf("error in Deserialize_int32. Wrong number of elements in byte array")
	} else {
		return int32(binary.LittleEndian.Uint32(encoded_int32)), nil
	}
}

func Deserialize_float(encoded_float32 []byte) (float32, error) {
	if len(encoded_float32) != int(types.JavonetFloatSize) {
		return 0, fmt.Errorf("error in Deserialize_float32. Wrong number of elements in byte array")
	} else {
		return float32(math.Float32frombits(binary.LittleEndian.Uint32(encoded_float32))), nil
	}
}

func Deserialize_byte(encoded_byte []byte) (byte, error) {
	if len(encoded_byte) != types.JavonetByteSize {
		return 0, fmt.Errorf("error in Deserialize_byte. Wrong number of elements in byte array")
	} else {
		return encoded_byte[0], nil
	}
}

func Deserialize_char(encoded_char []byte) (int8, error) {
	if len(encoded_char) != types.JavonetCharSize {
		return 0, fmt.Errorf("error in Deserialize_char. Wrong number of elements in byte array")
	} else {
		return int8(encoded_char[0]), nil
	}
}

func Deserialize_longlong(encoded_int64 []byte) (int64, error) {
	if len(encoded_int64) != types.JavonetLongLongSize {
		return 0, fmt.Errorf("error in Deserialize_int64. Wrong number of elements in byte array")
	} else {
		return int64(binary.LittleEndian.Uint64(encoded_int64)), nil
	}
}

func Deserialize_double(encoded_float64 []byte) (float64, error) {
	if len(encoded_float64) != int(types.JavonetDoubleSize) {
		return 0, fmt.Errorf("error in Deserialize_float64. Wrong number of elements in byte array")
	} else {
		return float64(math.Float64frombits(binary.LittleEndian.Uint64(encoded_float64))), nil
	}
}

func Deserialize_ulonglong(encoded_uint64 []byte) (uint64, error) {
	if len(encoded_uint64) != types.JavonetUnsignedLongLongSize {
		return 0, fmt.Errorf("error in Deserialize_uint64. Wrong number of elements in byte array")
	} else {
		return uint64(binary.LittleEndian.Uint64(encoded_uint64)), nil
	}
}

func Deserialize_uint(encoded_uint32 []byte) (uint32, error) {
	if len(encoded_uint32) != types.JavonetUnsignedIntegerSize {
		return 0, fmt.Errorf("error in Deserialize_uint32. Wrong number of elements in byte array")
	} else {
		return uint32(binary.LittleEndian.Uint32(encoded_uint32)), nil
	}
}
