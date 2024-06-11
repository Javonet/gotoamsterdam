package runtimefactory

import (
	"javonet.com/javonet/sdk/runtimecontext"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/runtimename"
)

// RuntimeFactory provides methods for creating runtime contexts.
// Each method corresponds to a specific runtime (CLR, JVM, .NET Core, Perl, Ruby, Node.js, Python) and returns a RuntimeContext instance for that runtime.
type RuntimeFactory struct {
	connectionType byte
	connectionData connectiondata.TcpConnectionData
}

// NewRuntimeFactory creates a new instance of RuntimeFactory with the provided connection type and data.
func NewRuntimeFactory(connType byte, connData connectiondata.TcpConnectionData) *RuntimeFactory {
	return &RuntimeFactory{connectionType: connType, connectionData: connData}
}

// Clr creates a RuntimeContext instance to interact with CLR runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Clr() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Clr, rf.connectionType, rf.connectionData)
}

// Jvm creates a RuntimeContext instance to interact with JVM runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Jvm() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Jvm, rf.connectionType, rf.connectionData)
}

// Netcore creates a RuntimeContext instance to interact with .NET Core runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Netcore() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Netcore, rf.connectionType, rf.connectionData)
}

// Perl creates a RuntimeContext instance to interact with Perl runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Perl() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Perl, rf.connectionType, rf.connectionData)
}

// Python creates a RuntimeContext instance to interact with Python runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Python() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Python, rf.connectionType, rf.connectionData)
}

// Ruby creates a RuntimeContext instance to interact with Ruby runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Ruby() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Ruby, rf.connectionType, rf.connectionData)
}

// Nodejs creates a RuntimeContext instance to interact with Node.js runtime.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (rf *RuntimeFactory) Nodejs() (*runtimecontext.RuntimeContext, error) {
	return runtimecontext.GetInstance(runtimename.Nodejs, rf.connectionType, rf.connectionData)
}
