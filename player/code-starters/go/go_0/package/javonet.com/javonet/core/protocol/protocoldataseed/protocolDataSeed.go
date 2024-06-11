package protocoldataseed

import (
	"math"

	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/runtimename"
)

var PythonCommandByteArray = []byte{5, 0,
	0, 0, 0, 0, 0, 0, 0,
	1, 3, 0, 2, 0, 0, 0, 5, 2, 0, 2, 0, 0, 0, 5, 6, 0, 1, 0, 0, 0, 5, 8, 1, 1, 8, 0, 0, 0, 100, 97, 116, 101, 116, 105, 109, 101, 1, 1, 4, 0, 0, 0, 100, 97, 116, 101, 1, 1, 5, 0, 0, 0, 116, 111, 100, 97, 121, 1, 1, 4, 0, 0, 0, 121, 101, 97, 114}

var payload_get_module = []interface{}{"datetime"}
var command_get_module = &command.Command{
	TargetRuntime: runtimename.Python,
	CommandType:   commandtype.GetModule,
	Payload:       payload_get_module,
}
var payload_get_type = []interface{}{command_get_module, "date"}
var command_get_type = &command.Command{
	TargetRuntime: runtimename.Python,
	CommandType:   commandtype.GetType,
	Payload:       payload_get_type,
}
var payload_invoke_static_method = []interface{}{command_get_type, "today"}
var command_invoke_static_method = &command.Command{
	TargetRuntime: runtimename.Python,
	CommandType:   commandtype.InvokeStaticMethod,
	Payload:       payload_invoke_static_method,
}
var payload_get_static_field = []interface{}{command_invoke_static_method, "year"}
var Command_get_static_field = &command.Command{
	TargetRuntime: runtimename.Python,
	CommandType:   commandtype.GetStaticField,
	Payload:       payload_get_static_field,
}

var JAVONET_string_serialized = []byte{1, 1, 7, 0, 0, 0, 74, 65, 86, 79, 78, 69, 84}
var NonAscii_string_serialized = []byte{1, 1, 14, 0, 0, 0, 197, 164, 194, 191, 207, 187, 195, 144, 195, 159, 196, 166, 197, 129}
var Empty_string_serialized = []byte{1, 1, 0, 0, 0, 0}
var Int_serialized =            []byte{2, 4, 89, 8, 0, 0}
var Bool_serialized = []byte{3, 1, 1}
var Float_serialized = []byte{4, 4, 195, 245, 170, 193}
var Byte_serialized = []byte{5, 1, 90}
var Char_serialized = []byte{6, 1, 91}
var Longlong_serialized = []byte{7, 8, 21, 205, 91, 7, 0, 0, 192, 255}
var Double_serialized = []byte{8, 8, 184, 86, 14, 60, 221, 154, 239, 63}
var Ullong_serialized = []byte{9, 8, 255, 255, 255, 255, 255, 255, 255, 255}
var Uint_serialized = []byte{10, 4, 254, 255, 255, 255}

var JAVONET_string_deserialized string = "JAVONET"
var NonAscii_string_deserialized string = "Ť¿ϻÐßĦŁ"
var Empty_string_deserialized string = ""
var Int_deserialized int32 = 2137
var Bool_deserialized bool = true
var Float_deserialized float32 = -21.37
var Byte_deserialized byte = 90
var Char_deserialized int8 = 91
var Longlong_deserialized int64 = -18014398386025195
var Double_deserialized float64 = 0.987654321
var Ullong_deserialized uint64 = math.MaxUint64
var Uint_deserialized uint32 = math.MaxUint32 - 1
