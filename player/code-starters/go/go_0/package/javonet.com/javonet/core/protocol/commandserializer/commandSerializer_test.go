//go:build unitTests
// +build unitTests

package commandserializer

import (
	"bytes"
	"testing"

	"javonet.com/javonet/core/protocol/protocoldataseed"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
)

func TestSerializeCommand_NestedCommand_CorrectByteArray(t *testing.T) {
	serializer := CommandSerializer{}

	response, err := serializer.Serialize(protocoldataseed.Command_get_static_field, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0), 0)
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !bytes.Equal(protocoldataseed.PythonCommandByteArray, response) {
		t.Fatal(t.Name() + " failed. Byte arrays do not match")
	}
}

func TestSerializeCommand_NilCommand_Error(t *testing.T) {
	serializer := NewCommandSerializer()
	_, err := serializer.Serialize(nil, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0), 0)
	if err == nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
}
