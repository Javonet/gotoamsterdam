//go:build unitTests
// +build unitTests

package connectiondata

import (
	"reflect"
	"testing"
)

func TestTcpConnectionData_Equal(t *testing.T) {
	tcpConnectionData1 := NewTcpConnectionData("127.0.0.1", 8080)
	tcpConnectionData2 := NewTcpConnectionData("127.0.0.1", 8080)
	if !reflect.DeepEqual(tcpConnectionData1, tcpConnectionData2) {
		t.Errorf("TcpConnectionData objects are not equal")
	}
}

func TestTcpConnectionData_NotEqual(t *testing.T) {
	tcpConnectionData1 := NewTcpConnectionData("127.0.0.1", 8080)
	tcpConnectionData2 := NewTcpConnectionData("127.0.0.1", 8081)
	if reflect.DeepEqual(tcpConnectionData1, tcpConnectionData2) {
		t.Errorf("TcpConnectionData objects are equal")
	}
}

func TestTcpConnectionData_GetAddresBytes(t *testing.T) {
	hostname := "localhost"
	port := uint16(8080)
	tcpConnectionData := NewTcpConnectionData(hostname, port)
	result := tcpConnectionData.GetAddressBytes()
	expected := []byte{127, 0, 0, 1}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Address bytes are not equal")
	}
}

func TestTcpConnectionData_GetAddresBytes2(t *testing.T) {
	hostname := "0.0.0.0"
	port := uint16(8080)
	tcpConnectionData := NewTcpConnectionData(hostname, port)
	result := tcpConnectionData.GetAddressBytes()
	expected := []byte{0, 0, 0, 0}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Address bytes are not equal")
	}
}

func TestTcpConnectionData_GetPortBytes(t *testing.T) {
	hostname := "localhost"
	port := uint16(8080)
	tcpConnectionData := NewTcpConnectionData(hostname, port)
	result := tcpConnectionData.GetPortBytes()
	expected := []byte{144, 31}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Port bytes are not equal")
	}
}
