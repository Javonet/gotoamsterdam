//go:build functionalTests
// +build functionalTests

package jsonfileresolver

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	dir               = filepath.Dir(filename)
	javonetSrcRoot    = filepath.Join(dir, "..", "..", "..", "..", "..", "..", "..")
	configPath        = filepath.Join(javonetSrcRoot, "/testResources/configuration-file/javonetconf.json")
)

func TestJsonFileResolver_GetLicenseKey(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	licenseKey, err := jfr.GetLicenseKey()
	if err != nil {
		t.Fatal(err)
	}
	if licenseKey != "your-license-key" {
		t.Errorf("Expected 'your-license-key', got %s", licenseKey)
	}
}

func TestJsonFileResolver_GetChannelType_NetcoreDefault(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	channelType, err := jfr.GetChannelType("netcore", "default")
	if err != nil {
		t.Fatal(err)
	}
	if channelType != "tcp" {
		t.Errorf("Expected 'tcp', got %s", channelType)
	}
}

func TestJsonFileResolver_GetChannelHost_NetcoreDefault(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	channelHost, err := jfr.GetChannelHost("netcore", "default")
	if err != nil {
		t.Fatal(err)
	}
	if channelHost != "127.0.0.1" {
		t.Errorf("Expected '127.0.0.1', got %s", channelHost)
	}
}

func TestJsonFileResolver_GetChannelPort_NetcoreCustom1(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	channelPort, err := jfr.GetChannelPort("netcore", "custom1")
	if err != nil {
		t.Fatal(err)
	}
	if channelPort != 8081 {
		t.Errorf("Expected 8081, got %d", channelPort)
	}
}

func TestJsonFileResolver_GetChannelType_PerlDefault(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	channelType, err := jfr.GetChannelType("perl", "default")
	if err != nil {
		t.Fatal(err)
	}
	if channelType != "inMemory" {
		t.Errorf("Expected 'inMemory', got %s", channelType)
	}
}

func TestGetChannelType_RubyCustom2_ThrowsException(t *testing.T) {
	jfr, err := NewJsonFileResolver(configPath)
	if err != nil {
		t.Fatal(err)
	}
	_, err = jfr.GetChannelType("ruby", "custom2")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
	expectedError := "Runtime config custom2 not found in configuration file for runtime ruby. Please check your configuration file."
	if err.Error() != expectedError {
		t.Errorf("Expected '%s', got '%s'", expectedError, err.Error())
	}
}
