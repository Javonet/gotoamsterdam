package invocationcontext

import (
	"javonet.com/javonet/core/interpreter"
	"javonet.com/javonet/utils/command"
	"javonet.com/javonet/utils/commandtype"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/exception"
)

// InvocationContext is a class that represents a context for invoking commands.
// It implements several interfaces for different types of interactions.
// This struct is used to construct chains of invocations, representing expressions of interaction that have not yet been executed.
// See: https://www.javonet.com/guides/v2/golang/foundations/invocation-context
type InvocationContext struct {
	isExecuted      bool
	runtimeName     byte
	connectionType  byte
	connectionData  connectiondata.TcpConnectionData
	currentCommand  *command.Command
	responseCommand *command.Command
	interpreter     *interpreter.Interpreter
}

func NewInvocationContext(runtimeName byte, connType byte, connData connectiondata.TcpConnectionData, currCmd *command.Command, isExec bool) *InvocationContext {
	return &InvocationContext{
		isExecuted:      isExec,
		runtimeName:     runtimeName,
		connectionType:  connType,
		connectionData:  connData,
		currentCommand:  currCmd,
		responseCommand: nil,
		interpreter:     interpreter.NewInterpreter(),
	}
}

// func (invCtx *InvocationContext) Close() {
// 	if invCtx.currentCommand.CommandType == commandtype.Reference {
// 		invCtx.currentCommand = command.NewCommand(invCtx.runtimeName, commandtype.DestructReference, &invCtx.currentCommand.Payload)
// 		invCtx.Execute()
// 	}
// }

func (invCtx *InvocationContext) GetCurrentCommand() *command.Command {
	return invCtx.currentCommand
}

// Execute executes the current command.
// Because invocation context is building the intent of executing particular expression on target environment, we call the initial state of invocation context as non-materialized.
// The non-materialized context wraps either single command or chain of recursively nested commands.
// Commands are becoming nested through each invocation of methods on Invocation Context.
// Each invocation triggers the creation of new Invocation Context instance wrapping the current command with new parent command valid for invoked method.
// Developer can decide on any moment of the materialization for the context taking full control of the chunks of the expression being transferred and processed on target runtime.
// See: https://www.javonet.com/guides/v2/golang/foundations/execute-method
func (invCtx *InvocationContext) Execute() (*InvocationContext, error) {
	var err error
	invCtx.responseCommand, err = invCtx.interpreter.Execute(invCtx.currentCommand, invCtx.connectionType, invCtx.connectionData)
	if err != nil {
		return nil, err
	}

	if invCtx.responseCommand.CommandType == commandtype.Exception {
		return nil, exception.ThrowException(invCtx.responseCommand)
	}

	if invCtx.currentCommand.CommandType == commandtype.CreateClassInstance {
		invCtx.currentCommand = invCtx.responseCommand
		invCtx.isExecuted = true
		return invCtx, nil
	}

	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.responseCommand, true), nil
}

// InvokeStaticMethod invokes a static method on the target runtime.
// methodName - The name of the method to invoke.
// args - The arguments to pass to the static method.
// Returns: A new InvocationContext instance that wraps the command to invoke the static method.
// See: https://www.javonet.com/guides/v2/golang/calling-methods/invoking-static-method
func (invCtx *InvocationContext) InvokeStaticMethod(methodName string, args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.InvokeStaticMethod, invCtx.createArgsArray(methodName, args))
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// InvokeInstanceMethod invokes an instance method on the target runtime.
// methodName - The name of the method to invoke.
// args - The arguments to pass to the instance method.
// Returns: A new InvocationContext instance that wraps the command to invoke the instance method.
// See: https://www.javonet.com/guides/v2/golang/calling-methods/creating-instance-and-calling-instance-methods
func (invCtx *InvocationContext) InvokeInstanceMethod(methodName string, args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.InvokeInstanceMethod, invCtx.createArgsArray(methodName, args))
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetStaticField retrieves the value of a static field from the target runtime.
// fieldName - The name of the field to retrieve.
// Returns: A new InvocationContext instance that wraps the command to get the static field.
// See: https://www.javonet.com/guides/v2/golang/fields-and-properties/getting-and-setting-values-for-static-fields-and-properties
func (invCtx *InvocationContext) GetStaticField(fieldName string) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.GetStaticField, &[]interface{}{fieldName})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// CreateInstance creates a new instance of a class in the target runtime.
// args - constructor arguments
// Returns: A new InvocationContext instance that wraps the command to create the instance.
// See: https://www.javonet.com/guides/v2/golang/calling-methods/creating-instance-and-calling-instance-methods
func (invCtx *InvocationContext) CreateInstance(args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.CreateClassInstance, &args)
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// SetStaticField sets the value of a static field in the target runtime.
// fieldName - The name of the field to set.
// value - The new value for the field.
// Returns: A new InvocationContext instance that wraps the command to set the static field.
// See: https://www.javonet.com/guides/v2/golang/fields-and-properties/getting-and-setting-values-for-static-fields-and-properties
func (invCtx *InvocationContext) SetStaticField(fieldName string, value interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.SetStaticField, &[]interface{}{fieldName, value})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetInstanceField retrieves the value of an instance field from the target runtime.
// fieldName - The name of the field to retrieve.
// Returns: A new InvocationContext instance that wraps the command to get the instance field.
// See: https://www.javonet.com/guides/v2/golang/fields-and-properties/getting-and- setting-values-for-instance-fields-and-properties
func (invCtx *InvocationContext) GetInstanceField(fieldName string) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.GetInstanceField, &[]interface{}{fieldName})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// SetInstanceField sets the value of an instance field in the target runtime.
// fieldName - The name of the field to set.
// value - The new value for the field.
// Returns: A new InvocationContext instance that wraps the command to set the instance field.
// See: https://www.javonet.com/guides/v2/golang/fields-and-properties/getting-and-setting-values-for-instance-fields-and-properties
func (invCtx *InvocationContext) SetInstanceField(fieldName string, value interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.SetInstanceField, &[]interface{}{fieldName, value})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetIndex retrieves the value at a specific index in an array from the target runtime.
// args - indexes of the array
// Returns: A new InvocationContext instance that wraps the command to get the index.
// See: https://www.javonet.com/guides/v2/golang/arrays-and-collections/one-dimensional-arrays
func (invCtx *InvocationContext) GetIndex(args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.ArrayGetItem, &args)
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetSize retrieves the size of an array from the target runtime.
// Returns: A new InvocationContext instance that wraps the command to get the size.
// See: https://www.javonet.com/guides/v2/golang/arrays-and-collections/one-dimensional-arrays
func (invCtx *InvocationContext) GetSize() *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.ArrayGetSize, &[]interface{}{})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetRank retrieves the rank of an array from the target runtime.
// Returns: A new InvocationContext instance that wraps the command to get the rank.
// See: https://www.javonet.com/guides/v2/golang/arrays-and-collections/one-dimensional-arrays
func (invCtx *InvocationContext) GetRank() *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.ArrayGetRank, &[]interface{}{})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// SetIndex sets the value at a specific index in an array in the target runtime.
// indexes - The indexes of the array.
// value - The new value for the index.
// Returns: A new InvocationContext instance that wraps the command to set the index.
// See: https://www.javonet.com/guides/v2/golang/arrays-and-collections/one-dimensional-arrays
func (invCtx *InvocationContext) SetIndex(indexes interface{}, value interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.ArraySetItem, &[]interface{}{indexes, value})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// InvokeGenericStaticMethod invokes a generic static method on the target runtime.
// methodName - The name of the method to invoke.
// args - The arguments to pass to the generic static method.
// Returns: A new InvocationContext instance that wraps the command to invoke the generic static method.
// See: https://www.javonet.com/guides/v2/golang/generics/calling-generic-static-method
func (invCtx *InvocationContext) InvokeGenericStaticMethod(methodName string, args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.InvokeGenericStaticMethod, invCtx.createArgsArray(methodName, args))
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// InvokeGenericMethod invokes a generic method on the target runtime.
// methodName - The name of the method to invoke.
// args - The arguments to pass to the generic method.
// Returns: A new InvocationContext instance that wraps the command to invoke the generic method.
// See: https://www.javonet.com/guides/v2/golang/generics/calling-generic-instance-method
func (invCtx *InvocationContext) InvokeGenericMethod(methodName string, args ...interface{}) *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.InvokeGenericMethod, invCtx.createArgsArray(methodName, args))
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetEnumName retrieves the name of an enum from the target runtime.
// Returns: A new InvocationContext instance that wraps the command to get the enum name.
// See: https://www.javonet.com/guides/v2/golang/enums/using-enum-type
func (invCtx *InvocationContext) GetEnumName() *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.GetEnumName, &[]interface{}{})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetEnumValue retrieves the value of an enum from the target runtime.
// Returns: A new InvocationContext instance that wraps the command to get the enum value.
// See: https://www.javonet.com/guides/v2/golang/enums/using-enum-type
func (invCtx *InvocationContext) GetEnumValue() *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.GetEnumValue, &[]interface{}{})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// GetRefValue retrieves the value of a reference from the target runtime.
// Returns: A new InvocationContext instance that wraps the command to get the reference value.
// See: https://www.javonet.com/guides/v2/golang/methods-arguments/passing-arguments-by-reference-with-ref-keyword
func (invCtx *InvocationContext) GetRefValue() *InvocationContext {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.GetRefValue, &[]interface{}{})
	return NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
}

// RetrieveArray retrieves an array from the target runtime.
// Returns: The array retrieved from the target runtime.
// See: https://www.javonet.com/guides/v2/golang/arrays-and-collections/collections
func (invCtx *InvocationContext) RetrieveArray() ([]interface{}, error) {
	invCtx.isExecuted = false
	localCommand := command.NewCommand(invCtx.runtimeName, commandtype.RetrieveArray, &[]interface{}{})
	localInvCtx := NewInvocationContext(invCtx.runtimeName, invCtx.connectionType, invCtx.connectionData, invCtx.buildCommand(localCommand), false)
	_, err := localInvCtx.Execute()
	return localInvCtx.responseCommand.Payload, err
}

// GetValue retrieves the value of the current command from the target runtime.
// The value of the current command.
// See: https://www.javonet.com/guides/v2/golang/foundations/execute-method
func (invCtx *InvocationContext) GetValue() interface{} {
	return invCtx.currentCommand.Payload[0]
}

func (invCtx *InvocationContext) buildCommand(cmd *command.Command) *command.Command {
	for i, _ := range cmd.Payload {
		cmd.Payload[i] = invCtx.encapsulatePayloadItem(cmd.Payload[i])
	}
	return cmd.PrependArgToPayload(&[]interface{}{invCtx.currentCommand})
}

func (invCtx *InvocationContext) encapsulatePayloadItem(payloadItem interface{}) *command.Command {
	switch element := payloadItem.(type) {
	case command.Command:
		for i, _ := range element.Payload {
			element.Payload[i] = invCtx.encapsulatePayloadItem(element.Payload[i])
		}
		return &element
	case (*InvocationContext):
		return element.GetCurrentCommand()
	case []interface{}:
		objectArray := make([]interface{}, len(element))
		for i, _ := range element {
			objectArray[i] = invCtx.encapsulatePayloadItem(element[i])
		}
		return command.NewCommand(invCtx.runtimeName, commandtype.Array, &objectArray)
	case interface{}:
		return command.NewCommand(invCtx.runtimeName, commandtype.Value, &[]interface{}{element})
	default:
		return nil
	}
}

func (rtmCtx *InvocationContext) createArgsArray(arg1 string, args []interface{}) *[]interface{} {
	argsArray := make([]interface{}, 1+len(args))
	argsArray[0] = arg1
	copy(argsArray[1:], args)
	return &argsArray
}
