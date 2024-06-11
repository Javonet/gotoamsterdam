package exception

import (
	"fmt"
	"strconv"
	"strings"

	"javonet.com/javonet/utils/command"
)

type ExceptionThrower struct{}

func ThrowException(stackCommand *command.Command) error {
	var builder strings.Builder
	builder.WriteString(stackCommand.Payload[2].(string))
	builder.WriteString("\n")
	builder.WriteString(stackCommand.Payload[3].(string))
	builder.WriteString("\n")
	stackTrace := string(getLocalStackTrace(stackCommand.Payload[4].(string), stackCommand.Payload[5].(string), stackCommand.Payload[6].(string), stackCommand.Payload[7].(string)))
	builder.WriteString(stackTrace)
	exceptionMessage := builder.String()
	return fmt.Errorf(exceptionMessage)
}

func getLocalStackTrace(stackTraceClasses string, stackTraceMethods string, stackTraceLines string, stackTraceFiles string) string {
	stackClasses := strings.Split(stackTraceClasses, "|")
	stackMethods := strings.Split(stackTraceMethods, "|")
	stackLines := strings.Split(stackTraceLines, "|")
	stackFiles := strings.Split(stackTraceFiles, "|")
	var builder strings.Builder
	for i := 0; i < len(stackLines); i++ {
		lineNumber, _ := strconv.Atoi(stackLines[i])
		builder.WriteString(fmt.Sprintf("%s(%s:%d) %s\n", stackMethods[i], stackFiles[i], lineNumber, stackClasses[i]))
	}
	return builder.String()
}
