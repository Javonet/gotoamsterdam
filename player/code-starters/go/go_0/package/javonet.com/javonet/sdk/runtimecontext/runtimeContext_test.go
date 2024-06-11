//go:build functionalTests
// +build functionalTests

package runtimecontext

import (
	"testing"

	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
	"javonet.com/javonet/utils/runtimename"
)

func TestRuntimeContext_InvokeMultipleRuntimeContexts_AreEqual(t *testing.T) {
	rc1, err := GetInstance(runtimename.Perl, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0))
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := GetInstance(runtimename.Perl, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0))
	if err != nil {
		t.Fatal(err)
	}
	if rc1 != rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeContext_InvokeMultipleRuntimeContexts_AreEqual2(t *testing.T) {
	rc1, err := GetInstance(runtimename.Ruby, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := GetInstance(runtimename.Ruby, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	if rc1 != rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeContext_InvokeMultipleRuntimeContexts_AreDifferent1(t *testing.T) {
	rc1, err := GetInstance(runtimename.Jvm, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := GetInstance(runtimename.Jvm, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 8080))
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeContext_InvokeMultipleRuntimeContexts_AreDifferent2(t *testing.T) {
	rc1, err := GetInstance(runtimename.Netcore, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := GetInstance(runtimename.Netcore, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.2", 8080))
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestRuntimeContext_InvokeMultipleRuntimeContexts_AreDifferent3(t *testing.T) {
	rc1, err := GetInstance(runtimename.Jvm, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := GetInstance(runtimename.Netcore, connectiontype.Tcp, *connectiondata.NewTcpConnectionData("127.0.0.1", 8080))
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}
