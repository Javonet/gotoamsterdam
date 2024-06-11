//go:build functionalTests
// +build functionalTests

package configruntimefactory

import (
	"path/filepath"
	"runtime"
	"testing"
)

var (
	_, filename, _, _ = runtime.Caller(0)
	dir               = filepath.Dir(filename)
	javonetSrcRoot    = filepath.Join(dir, "..", "..", "..", "..", "..", "..")
	configPath        = filepath.Join(javonetSrcRoot, "/testResources/configuration-file/javonetconf.json")
)

func TestConfigRuntimeFactory_InvokeMultipleRuntimeContexts_AreEqual(t *testing.T) {
	rc1, err := NewConfigRuntimeFactory(configPath).Nodejs("")
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewConfigRuntimeFactory(configPath).Nodejs("")
	if err != nil {
		t.Fatal(err)
	}
	if rc1 != rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestConfigRuntimeFactory_InvokeMultipleRuntimeContexts_AreDifferent(t *testing.T) {
	rc1, err := NewConfigRuntimeFactory(configPath).Ruby("")
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewConfigRuntimeFactory(configPath).Python("")
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}

func TestConfigRuntimeFactory_InvokeMultipleRuntimeContexts_AreDifferent2(t *testing.T) {
	rc1, err := NewConfigRuntimeFactory(configPath).Netcore("")
	if err != nil {
		t.Fatal(err)
	}
	rc2, err := NewConfigRuntimeFactory(configPath).Netcore("custom1")
	if err != nil {
		t.Fatal(err)
	}
	if rc1 == rc2 {
		t.Fatal(t.Name() + " failed")
	}
}
