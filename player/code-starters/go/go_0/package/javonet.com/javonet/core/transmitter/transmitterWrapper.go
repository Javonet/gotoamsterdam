package transmitter

/*
#cgo linux LDFLAGS: -ldl
#cgo darwin LDFLAGS: -ldl
#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#ifdef _WIN32
#include <windows.h>
#include <direct.h>
HINSTANCE goLibInstance = NULL;
#endif //_WIN32
#if defined(__linux__) || defined(__APPLE__)
#include <dlfcn.h> //dlsym etc.
#include <unistd.h>
#include <stdlib.h>
void* goLibInstance = NULL;
#endif // __linux__ || __APPLE__

typedef unsigned char byte;
typedef int (*Func1)(byte*, int);
typedef int (*Func2)(const char*, const char*, const char*, const char*);
typedef const char* (*Func3)();
typedef int (*Func4)(const char*);

Func1 SendCommand_fn;
Func1 ReadResponse_fn;
Func2 Activate_fn;
Func3 GetNativeError_fn;
Func4 SetConfigSource_Fn;

int ImportLibraryAndExportedFunctions(const char* path) {
	if(!goLibInstance) {
		char fullPath[1024];
#ifdef _WIN32
		sprintf(fullPath, "%s%s", path, "\\Binaries\\Native\\Windows\\X64\\JavonetGoRuntimeNative.dll");
		if (GetFileAttributesA(fullPath) == INVALID_FILE_ATTRIBUTES) return 1;

		goLibInstance = LoadLibraryA(fullPath);
		if (!goLibInstance) return 2;
		SendCommand_fn = (Func1)GetProcAddress(goLibInstance, "SendCommand");
		ReadResponse_fn = (Func1)GetProcAddress(goLibInstance, "ReadResponse");
		Activate_fn = (Func2)GetProcAddress(goLibInstance, "Activate");
		GetNativeError_fn = (Func3)GetProcAddress(goLibInstance, "GetNativeError");
		SetConfigSource_Fn = (Func4)GetProcAddress(goLibInstance, "SetConfigSource");
#endif //_WIN32
#ifdef __linux__
		sprintf(fullPath, "%s%s", path, "/Binaries/Native/Linux/X64/libJavonetGoRuntimeNative.so");
#endif	//__linux__
#ifdef  __APPLE__
		sprintf(fullPath, "%s%s", path, "/Binaries/Native/MacOs/X64/libJavonetGoRuntimeNative.dylib");
#endif	// __APPLE__

#if defined(__linux__) || defined(__APPLE__)
		goLibInstance = dlopen(fullPath, RTLD_LAZY);
		char* error = dlerror();
		if (goLibInstance == NULL || error != NULL) return 3;
		SendCommand_fn = (Func1)dlsym(goLibInstance, "SendCommand");
		ReadResponse_fn = (Func1)dlsym(goLibInstance, "ReadResponse");
		Activate_fn = (Func2)dlsym(goLibInstance, "Activate");
		GetNativeError_fn = (Func3)(Func2)dlsym(goLibInstance, "GetNativeError");
		SetConfigSource_Fn = (Func4)dlsym(goLibInstance, "SetConfigSource");
#endif	//__linux__ || __APPLE__
	}
	if(!SendCommand_fn || !ReadResponse_fn || !Activate_fn || !GetNativeError_fn || !SetConfigSource_Fn)
	{
		return 4;
	}
	return 0;
}

int SendCommand(byte* messageByteArray, int messageByteArrayLen) {
	return SendCommand_fn(messageByteArray, messageByteArrayLen);
}

int ReadResponse(byte* responseByteArray, int responseByteArrayLen) {
	return ReadResponse_fn(responseByteArray, responseByteArrayLen);
}

int Activate(const char* licenseKey, const char* proxyHost, const char* proxyUserName, const char* proxyPassword) {
	return Activate_fn(licenseKey, proxyHost, proxyUserName, proxyPassword);
}

const char* GetNativeError() {
	return GetNativeError_fn();
}

int SetConfigSource(const char* configSource) {
	return SetConfigSource_Fn(configSource);
}

*/
import "C"
import (
	"fmt"
	"os"
	"strconv"
	"unsafe"

	"javonet.com/javonet/utils/binariesextractor"
)

// TransmitterWrapper communicates with native channel
type TransmitterWrapper struct {
}

func init() {
	// Extract Native binaries from zip
	var be = binariesextractor.NewBinariesExtractor()
	err := be.ExtractFromBinariesZip("Native")
	if err != nil {
		panic("Failed to extract binaries: " + err.Error())
	}

	// Load native library
	binariesPath, _ := os.Getwd()
	binariesPathC := C.CString(binariesPath)
	defer C.free(unsafe.Pointer(binariesPathC))
	resultC := C.ImportLibraryAndExportedFunctions(binariesPathC)
	result := int(resultC)
	if result != 0 {
		panic("Failed to load native library: " + strconv.Itoa(result))
	}
}

// SendCommand sends command to native channel
func (tw *TransmitterWrapper) SendCommand(messageByteArray []byte) ([]byte, error) {
	//old way - do not delete, for reference
	//messageByteArrayC := (*C.uchar)(C.malloc(C.size_t(len(messageByteArray))))
	//defer C.free(unsafe.Pointer(messageByteArrayC))
	//messageByteArrayC = (*C.uchar)(unsafe.Pointer(&messageByteArray[0]))
	//responseByteArrayLen := int(C.SendCommand(messageByteArrayC, C.int(len(messageByteArray))))
	responseByteArrayLen := int(C.SendCommand((*C.byte)(&messageByteArray[0]), C.int(len(messageByteArray))))

	if responseByteArrayLen > 0 {
		//old way - do not delete, for reference
		//responseByteArrayC := (*C.uchar)(C.malloc(C.size_t(responseByteArrayLen)))
		//defer C.free(unsafe.Pointer(responseByteArrayC))
		//C.ReadResponse(responseByteArrayC, C.int(responseByteArrayLen))
		//responseByteArray := C.GoBytes(unsafe.Pointer(responseByteArrayC), C.int(responseByteArrayLen))
		responseByteArray := make([]byte, responseByteArrayLen)
		C.ReadResponse((*C.byte)(&responseByteArray[0]), C.int(responseByteArrayLen))
		return responseByteArray, nil
	} else if responseByteArrayLen == 0 {
		errorMessage := "response is empty"
		return nil, fmt.Errorf(errorMessage)
	} else {
		errorMessage := C.GoString(C.GetNativeError())
		return nil, fmt.Errorf("Javonet native error code: " + strconv.Itoa(responseByteArrayLen) + ". " + errorMessage)
	}
}

// Activate Javonet
func (tw *TransmitterWrapper) Activate(licenseKey string, proxyHost string, proxyUserName string, proxyPassword string) (int, error) {
	licenseKeyC := C.CString(licenseKey)
	proxyhostC := C.CString(proxyHost)
	proxyUserNameC := C.CString(proxyUserName)
	proxyPasswordC := C.CString(proxyPassword)
	defer C.free(unsafe.Pointer(licenseKeyC))
	defer C.free(unsafe.Pointer(proxyhostC))
	defer C.free(unsafe.Pointer(proxyUserNameC))
	defer C.free(unsafe.Pointer(proxyPasswordC))

	activationResult := int(C.Activate((*C.char)(licenseKeyC), (*C.char)(proxyhostC), (*C.char)(proxyUserNameC), (*C.char)(proxyPasswordC)))
	if activationResult < 0 {
		errorMessage := C.GoString(C.GetNativeError())
		return activationResult, fmt.Errorf("Javonet activation result: " + strconv.Itoa(activationResult) + ". " + errorMessage)
	} else {
		return activationResult, nil
	}
}

// SetConfigSource sets config source
func (tw *TransmitterWrapper) SetConfigSource(configPath string) (int, error) {
	configPathC := C.CString(configPath)
	defer C.free(unsafe.Pointer(configPathC))
	result := int(C.SetConfigSource((*C.char)(configPathC)))
	if result < 0 {
		errorMessage := C.GoString(C.GetNativeError())
		return result, fmt.Errorf("Javonet set config source result: " + strconv.Itoa(result) + ". " + errorMessage)
	} else {
		return result, nil
	}
}
