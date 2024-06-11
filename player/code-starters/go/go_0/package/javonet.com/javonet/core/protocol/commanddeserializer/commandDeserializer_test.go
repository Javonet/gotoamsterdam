//go:build unitTests
// +build unitTests

package commanddeserializer

import (
	"reflect"
	"testing"

	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/runtimename"
)

func TestDeserializeCommand_MessageWithCommand_CorrectCommand(t *testing.T) {
	command_get_module := command.NewCommand(runtimename.Python, commandtype.GetModule, &[]interface{}{"datetime"})
	command_get_type := command.NewCommand(runtimename.Python, commandtype.GetType, &[]interface{}{command_get_module, "date"})

	expectedResponse := command_get_type
	messageByteArray := []byte{5, 0, 0, 0, 0, 0, 0, 0, 0, 5, 6, 0, 1, 0, 0, 0, 5, 8, 1, 0, 8, 0, 0, 0, 100, 97, 116, 101, 116, 105, 109, 101, 1, 0, 4, 0, 0, 0, 100, 97, 116, 101}
	deserializer := NewCommandDeserializer(messageByteArray)
	response, err := deserializer.Deserialize()
	if err != nil {
		t.Fatal(t.Name() + " failed. " + err.Error())
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed. ")
	}
}
