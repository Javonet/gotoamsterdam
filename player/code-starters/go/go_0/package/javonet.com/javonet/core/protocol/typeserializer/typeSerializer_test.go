//go:build unitTests
// +build unitTests

package typeserializer

import (
	"bytes"
	"testing"

	"javonet.com/javonet/core/protocol/protocoldataseed"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/runtimename"
	"javonet.com/javonet/utils/types"
)

func Test_TypeSerializer_Command_SimpleCommand_CorrectByteArray(t *testing.T) {
	expectedResponse := []byte{types.JavonetCommand, 1, 0, 0, 0, 0, 1}
	payload := make([]interface{}, 1)
	payload[0] = "COMMAND"
	command := &command.Command{
		TargetRuntime: runtimename.Clr,
		CommandType:   types.JavonetString,
		Payload:       payload,
	}
	response, err := Serialize_command(command)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(expectedResponse, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_String_JAVONET(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.JAVONET_string_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.JAVONET_string_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_String_Nonasciistring(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.NonAscii_string_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.NonAscii_string_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_String_Emptystring(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Empty_string_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Empty_string_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}
func Test_TypeSerializer_Int(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Int_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Int_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Bool(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Bool_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Bool_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Float(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Float_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Float_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Byte(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Byte_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Byte_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Char(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Char_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Char_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Longlong(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Longlong_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Longlong_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Double(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Double_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Double_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Ullong(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Ullong_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Ullong_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func Test_TypeSerializer_Uint(t *testing.T) {
	response, err := Serialize_primitive(protocoldataseed.Uint_deserialized)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.Uint_serialized, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}
