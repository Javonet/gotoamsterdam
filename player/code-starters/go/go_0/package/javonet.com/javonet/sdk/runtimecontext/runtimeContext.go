package runtimecontext

import (
	"fmt"

	"javonet.com/javonet/core/interpreter"
	"javonet.com/javonet/sdk/invocationcontext"
	"javonet.com/javonet/utils/binariesextractor"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
	"javonet.com/javonet/utils/exception"
	"javonet.com/javonet/utils/runtimenamehandler"
)

// RuntimeContext represents a single context which allows interaction with a selected technology.
// It refers to a single instance of the called runtime within a particular target OS process.
// This can be either the local currently running process (inMemory) or a particular remote process identified by the IP Address and PORT of the target Javonet instance.
// Multiple Runtime Contexts can be initialized within one process.
// Calling the same technology on inMemory communication channel will return the existing instance of runtime context.
// Calling the same technology on TCP channel but on different nodes will result in unique Runtime Contexts.
// Within the runtime context, any number of libraries can be loaded and any objects from the target technology can be interacted with, as they are aware of each other due to sharing the same memory space and same runtime instance.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
type RuntimeContext struct {
	isExecuted      bool
	runtimeName     byte
	connectionType  byte
	connectionData  connectiondata.TcpConnectionData
	currentCommand  *command.Command
	responseCommand *command.Command
	interpreter     *interpreter.Interpreter
}

type networkRuntimeContextKey struct {
	runtime     byte
	connDataStr string
}

var memoryRuntimeContext = make(map[byte]*RuntimeContext)
var networkRuntimeContext = make(map[networkRuntimeContextKey]*RuntimeContext)

func newRuntimeContext(runtimeName byte, connType byte, connData connectiondata.TcpConnectionData) *RuntimeContext {
	// Extract binaries from the zip file
	be := binariesextractor.NewBinariesExtractor()
	runtimeNameStr, _ := runtimenamehandler.GetName(runtimeName)
	err := be.ExtractFromBinariesZip(runtimeNameStr)
	if err != nil {
		panic("Failed to extract binaries: " + err.Error())
	}

	// Create new runtime context
	return &RuntimeContext{
		isExecuted:      false,
		runtimeName:     runtimeName,
		connectionType:  connType,
		connectionData:  connData,
		currentCommand:  nil,
		responseCommand: nil,
		interpreter:     interpreter.NewInterpreter(),
	}
}

func GetInstance(runtimeName byte, connType byte, connData connectiondata.TcpConnectionData) (*RuntimeContext, error) {
	switch connType {
	case connectiontype.InMemory:
		if memoryRuntimeContext[runtimeName] == nil {
			runtimeContext := newRuntimeContext(runtimeName, connType, connData)
			memoryRuntimeContext[runtimeName] = runtimeContext
			return runtimeContext, nil
		} else {
			memoryRuntimeContext[runtimeName].currentCommand = nil
			return memoryRuntimeContext[runtimeName], nil
		}
	case connectiontype.Tcp:
		mapKey := networkRuntimeContextKey{
			runtime:     runtimeName,
			connDataStr: fmt.Sprintf("%s:%d", connData.IpAddress.String(), connData.Port),
		}
		if networkRuntimeContext[mapKey] == nil {
			runtimeContext := newRuntimeContext(runtimeName, connType, connData)
			networkRuntimeContext[mapKey] = runtimeContext
			return runtimeContext, nil
		} else {
			networkRuntimeContext[mapKey].currentCommand = nil
			return networkRuntimeContext[mapKey], nil
		}
	default:
		return nil, fmt.Errorf("invalid connection type")
	}
}

// Execute method sends the current command or chain of commands to the target runtime for execution.
// The initial state of RuntimeContext is non-materialized, wrapping either a single command or a chain of recursively nested commands.
// Commands become nested through each invocation of methods on RuntimeContext.
// Each invocation triggers the creation of a new RuntimeContext instance wrapping the current command with a new parent command.
// The developer can decide at any moment of the materialization for the context, taking full control of the chunks of the expression being transferred and processed on the target runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rtmCtx *RuntimeContext) Execute() error {
	var err error
	rtmCtx.responseCommand, err = rtmCtx.interpreter.Execute(rtmCtx.currentCommand, rtmCtx.connectionType, rtmCtx.connectionData)
	rtmCtx.currentCommand = nil
	if err != nil {
		return err
	}

	if rtmCtx.responseCommand.CommandType == commandtype.Exception {
		return exception.ThrowException(rtmCtx.responseCommand)
	}

	rtmCtx.isExecuted = true
	return nil
}

// LoadLibrary method allows you to reference and use modules or packages written in various languages.
// This method allows you to use any library from all supported technologies. The necessary libraries need to be referenced.
// If the library has dependencies on other libraries, the latter needs to be added first.
// After referencing the library, any objects stored in this package can be used. Use static classes, create instances, call methods, use fields and properties, and much more.
// libraryPath - relative or full path to the library.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/getting-started/adding-references-to-libraries
func (rtmCtx *RuntimeContext) LoadLibrary(libraryPath string) (*RuntimeContext, error) {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.LoadLibrary, &[]interface{}{libraryPath})
	rtmCtx.currentCommand = rtmCtx.buildCommand(localCommand)
	err := rtmCtx.Execute()
	if err != nil {
		return nil, err
	}
	return rtmCtx, nil
}

// GetType method retrieves a type from the called runtime.
// The type can be a class, interface or enum. The type can be retrieved from any referenced library.
// typeName - the name of the type to be retrieved.
// args - optional arguments.
func (rtmCtx *RuntimeContext) GetType(typeName string, args ...interface{}) *invocationcontext.InvocationContext {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.GetType, rtmCtx.createArgsArray(typeName, args))
	rtmCtx.currentCommand = nil
	return invocationcontext.NewInvocationContext(rtmCtx.runtimeName, rtmCtx.connectionType, rtmCtx.connectionData, rtmCtx.buildCommand(localCommand), false)
}

// Cast method allows you to cast values when working with methods from the called runtime that require specific types of arguments.
// The arguments include the target type and the value to be cast. The target type must be retrieved from the called runtime using the GetType method.
// After casting the value, it can be used as an argument when invoking methods.
// args - the target type and the value to be cast.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/casting/casting
func (rtmCtx *RuntimeContext) Cast(args ...interface{}) *invocationcontext.InvocationContext {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.Cast, &args)
	rtmCtx.currentCommand = nil
	return invocationcontext.NewInvocationContext(rtmCtx.runtimeName, rtmCtx.connectionType, rtmCtx.connectionData, rtmCtx.buildCommand(localCommand), false)
}

// GetEnumItem method allows you to retrieve an item from an enum type in the called runtime.
// The arguments include the enum type and the name of the item. The enum type must be retrieved from the called runtime using the GetType method.
// After retrieving the item, it can be used as an argument when invoking methods or for other operations.
// args - the enum type and the item name.
// See: https://www.javonet.com/guides/v2/golang/enums/using-enum-type
func (rtmCtx *RuntimeContext) GetEnumItem(args ...interface{}) *invocationcontext.InvocationContext {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.GetEnumItem, &args)
	rtmCtx.currentCommand = nil
	return invocationcontext.NewInvocationContext(rtmCtx.runtimeName, rtmCtx.connectionType, rtmCtx.connectionData, rtmCtx.buildCommand(localCommand), false)
}

// AsRef method allows you to create a reference when working with methods from the called runtime that require arguments to be passed by reference.
// The arguments include the value and optionally the type of the reference. The type must be retrieved from the called runtime using the GetType method.
// After creating the reference, it can be used as an argument when invoking methods.
// args - The value and optionally the type of the reference.
// See: https://www.javonet.com/guides/v2/golang/methods-arguments/passing-arguments-by-reference-with-ref-keyword
func (rtmCtx *RuntimeContext) AsRef(args ...interface{}) *invocationcontext.InvocationContext {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.AsRef, &args)
	rtmCtx.currentCommand = nil
	return invocationcontext.NewInvocationContext(rtmCtx.runtimeName, rtmCtx.connectionType, rtmCtx.connectionData, rtmCtx.buildCommand(localCommand), false)
}

// AsOut method allows you to create a reference when working with methods from the called runtime that require arguments to be passed by reference.
// The arguments include the value and optionally the type of the reference. The type must be retrieved from the called runtime using the GetType method.
// After creating the reference, it can be used as an argument when invoking methods.
// args - The value and optionally the type of the reference.
// See: https://www.javonet.com/guides/v2/golang/methods-arguments/passing-arguments-by-reference-with-out-keyword
func (rtmCtx *RuntimeContext) AsOut(args ...interface{}) *invocationcontext.InvocationContext {
	localCommand := command.NewCommand(rtmCtx.runtimeName, commandtype.AsOut, &args)
	rtmCtx.currentCommand = nil
	return invocationcontext.NewInvocationContext(rtmCtx.runtimeName, rtmCtx.connectionType, rtmCtx.connectionData, rtmCtx.buildCommand(localCommand), false)
}

func (rtmCtx *RuntimeContext) buildCommand(cmd *command.Command) *command.Command {
	for i, _ := range cmd.Payload {
		cmd.Payload[i] = rtmCtx.encapsulatePayloadItem(cmd.Payload[i])
	}
	if rtmCtx.currentCommand == nil {
		return cmd
	} else {
		return cmd.PrependArgToPayload(&[]interface{}{rtmCtx.currentCommand})
	}
}

func (rtmCtx *RuntimeContext) encapsulatePayloadItem(payloadItem interface{}) *command.Command {
	switch element := payloadItem.(type) {
	case command.Command:
		for i, _ := range element.Payload {
			element.Payload[i] = rtmCtx.encapsulatePayloadItem(element.Payload[i])
		}
		return &element
	case (*invocationcontext.InvocationContext):
		return element.GetCurrentCommand()
	case []interface{}:
		objectArray := make([]interface{}, len(element))
		for i, _ := range element {
			objectArray[i] = rtmCtx.encapsulatePayloadItem(element[i])
		}
		return command.NewCommand(rtmCtx.runtimeName, commandtype.Array, &objectArray)
	case interface{}:
		return command.NewCommand(rtmCtx.runtimeName, commandtype.Value, &[]interface{}{element})
	default:
		return nil
	}
}

func (rtmCtx *RuntimeContext) createArgsArray(arg1 string, args []interface{}) *[]interface{} {
	argsArray := make([]interface{}, 1+len(args))
	argsArray[0] = arg1
	copy(argsArray[1:], args)
	return &argsArray
}
