//go:build unitTests
// +build unitTests

package typedeserializer

import (
	"testing"

	"javonet.com/javonet/core/protocol/protocoldataseed"
	"javonet.com/javonet/utils/stringencodingmode"
)

func Test_TypeDeserializer_String_JAVONET(t *testing.T) {
	response, err := Deserialize_string(stringencodingmode.UTF8, protocoldataseed.JAVONET_string_serialized[6:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.JAVONET_string_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_String_Nonascii(t *testing.T) {
	response, err := Deserialize_string(stringencodingmode.UTF8, protocoldataseed.NonAscii_string_serialized[6:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.NonAscii_string_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_String_Empty(t *testing.T) {
	response, err := Deserialize_string(stringencodingmode.UTF8, protocoldataseed.Empty_string_serialized[6:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Empty_string_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Int(t *testing.T) {
	response, err := Deserialize_int(protocoldataseed.Int_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Int_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Bool(t *testing.T) {
	response, err := Deserialize_bool(protocoldataseed.Bool_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Bool_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Float(t *testing.T) {
	response, err := Deserialize_float(protocoldataseed.Float_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Float_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Byte(t *testing.T) {
	response, err := Deserialize_byte(protocoldataseed.Byte_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Byte_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Char(t *testing.T) {
	response, err := Deserialize_char(protocoldataseed.Char_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Char_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_LongLong(t *testing.T) {
	response, err := Deserialize_longlong(protocoldataseed.Longlong_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Longlong_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_Double(t *testing.T) {
	response, err := Deserialize_double(protocoldataseed.Double_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Double_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_UnsignedLongLong(t *testing.T) {
	response, err := Deserialize_ulonglong(protocoldataseed.Ullong_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Ullong_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeDeserializer_UnsignedInt(t *testing.T) {
	response, err := Deserialize_uint(protocoldataseed.Uint_serialized[2:])
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if protocoldataseed.Uint_deserialized != response {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}
