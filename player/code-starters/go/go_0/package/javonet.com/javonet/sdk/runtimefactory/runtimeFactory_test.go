//go:build functionalTests
// +build functionalTests

package runtimefactory

import (
	"testing"

	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
)

func TestRuntimeFactory_InvokeMultipleRuntimeContexts_AreEqual(t *testing.T) {
	rc1, err := NewRuntimeFactory(connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0)).Perl()
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewRuntimeFactory(connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0)).Perl()
	if err != nil {
		t.Fatal(err)
	}
	if rc1 != rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeFactory_InvokeMultipleRuntimeContexts_AreEqual2(t *testing.T) {
	rc1, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Netcore()
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Netcore()
	if err != nil {
		t.Fatal(err)
	}
	if rc1 != rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeFactory_InvokeMultipleRuntimeContexts_AreDifferent1(t *testing.T) {
	rc1, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Jvm()
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewRuntimeFactory(connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 8080)).Jvm()
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeFactory_InvokeMultipleRuntimeContexts_AreDifferent2(t *testing.T) {
	rc1, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Netcore()
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.2", 8080)).Netcore()
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeFactory_InvokeMultipleRuntimeContexts_AreDifferent3(t *testing.T) {
	rc1, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Jvm()
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewRuntimeFactory(connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080)).Netcore()
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}
