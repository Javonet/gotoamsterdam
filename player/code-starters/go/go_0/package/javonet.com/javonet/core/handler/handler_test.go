//go:build false
// +build false

package handler

import (
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"

	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/runtimename"
)

func changeDir() {
	_, filename, _, _ := runtime.Caller(0)
	var dir string
	if runtime.GOOS == "windows" {
		dir = path.Join(path.Dir(filename), "\\..\\..")
	} else {
		dir = path.Join(path.Dir(filename), "/../..")
	}
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}

func TestHandler_LoadLibraryCommand_Success(t *testing.T) {
	changeDir()
	cmd := command.NewCommand(runtimename.Go, commandtype.LoadLibrary, &[]interface{}{"plugin/goPlugin.so"})
	expectedResponse := command.NewCommand(runtimename.Go, commandtype.Value, &[]interface{}{true})
	hndl := NewHandler()
	response, err := hndl.HandleCommand(cmd)
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed: Wrong response")
	}
}

func TestHandler_LoadLibraryCommandWrong_Failure(t *testing.T) {
	changeDir()
	cmd := command.NewCommand(runtimename.Go, commandtype.LoadLibrary, &[]interface{}{"plugin/goPlugin.so", "dummy argument"})
	expectedResponse := command.NewCommand(runtimename.Go, commandtype.Exception, &[]interface{}{})
	hndl := NewHandler()
	response, err := hndl.HandleCommand(cmd)
	if err == nil {
		t.Fatal(t.Name() + " failed: ")
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed: Wrong response")
	}
}

func TestHandler_GetTypeCommand_Success(t *testing.T) {
	changeDir()
	cmd := command.NewCommand(runtimename.Go, commandtype.GetType, &[]interface{}{"plugin/goPlugin.so"})
	expectedResponse := command.NewCommand(runtimename.Go, commandtype.Value, &[]interface{}{true})
	hndl := NewHandler()
	response, err := hndl.HandleCommand(cmd)
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed: Wrong response")
	}
}

func TestHandler_GetTypeCommand_Failure(t *testing.T) {
	changeDir()
	cmd := command.NewCommand(runtimename.Go, commandtype.GetType, &[]interface{}{"plugin/goPluginwrong.so"})
	hndl := NewHandler()
	_, err := hndl.HandleCommand(cmd)
	if err == nil {
		t.Fatal(t.Name() + " failed: ")
	}
}

func TestHandler_InvokeStaticMethodCommand_NoArgs_Success(t *testing.T) {
	changeDir()
	hndl := NewHandler()
	cmd := command.NewCommand(runtimename.Go, commandtype.GetType, &[]interface{}{"plugin/goPlugin.so"})
	response, err := hndl.HandleCommand(cmd)
	expectedResponse := command.NewCommand(runtimename.Go, commandtype.Value, &[]interface{}{true})
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed: Wrong response in GetType")
	}

	cmd2 := command.NewCommand(runtimename.Go, commandtype.InvokeStaticMethod, &[]interface{}{"SampleFuncPrintHello"})
	response, err = hndl.HandleCommand(cmd2)
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if len(response.Payload[0].(string)) != 36 {
		t.Fatal(t.Name() + " failed: Wrong response in GetType")
	}
}

func TestHandler_InvokeStaticMethodCommand_WithArgs_Success(t *testing.T) {
	changeDir()
	hndl := NewHandler()
	cmd := command.NewCommand(runtimename.Go, commandtype.GetType, &[]interface{}{"plugin/goPlugin.so"})
	response, err := hndl.HandleCommand(cmd)
	expectedResponse := command.NewCommand(runtimename.Go, commandtype.Value, &[]interface{}{true})
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if !reflect.DeepEqual(*response, *expectedResponse) {
		t.Fatal(t.Name() + " failed: Wrong response in GetType")
	}

	cmd2 := command.NewCommand(runtimename.Go, commandtype.InvokeStaticMethod, &[]interface{}{"SampleFuncPrintArg", "A", "B"})
	response, err = hndl.HandleCommand(cmd2)
	if err != nil {
		t.Fatal(t.Name() + " failed: " + err.Error())
	}
	if len(response.Payload[0].(string)) != 36 {
		t.Fatal(t.Name() + " failed: Wrong response in GetType")
	}
}
