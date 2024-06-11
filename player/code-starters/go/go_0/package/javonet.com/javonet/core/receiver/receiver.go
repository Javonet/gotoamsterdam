package main

//#include <stdlib.h>
import "C"
import (
	"unsafe"

	"javonet.com/javonet/core/interpreter"
)

func main() {}

//export HeartBeat
func HeartBeat(messageByteArray *C.uchar, messageByteArrayLen C.int) (*C.uchar, C.int) {
	response := make([]byte, 2)
	response[0] = 49
	response[1] = 48
	responseByteArray := (*C.uchar)(C.CBytes(response))
	responseByteArrayLen := C.int(len(response))
	//defer C.free(unsafe.Pointer(responseByteArray))
	return responseByteArray, responseByteArrayLen
}

//export SendCommandReceiver
func SendCommandReceiver(messageByteArray *C.uchar, messageByteArrayLen C.int) (*C.uchar, C.int) {
	interpreter := interpreter.NewInterpreter()
	response, err := interpreter.Process(C.GoBytes(unsafe.Pointer(messageByteArray), C.int(messageByteArrayLen)))
	if err != nil {
		return nil, C.int(0)
	} else {
		responseByteArray := (*C.uchar)(C.CBytes(response))
		responseByteArrayLen := C.int(len(response))
		//defer C.free(unsafe.Pointer(responseByteArray))
		return responseByteArray, responseByteArrayLen
	}

}
