package configruntimefactory

import (
	"errors"

	"javonet.com/javonet/core/transmitter"
	"javonet.com/javonet/sdk/runtimecontext"
	"javonet.com/javonet/sdk/tools/jsonfileresolver"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
	"javonet.com/javonet/utils/runtimename"
	"javonet.com/javonet/utils/runtimenamehandler"
)

// ConfigRuntimeFactory provides methods for creating runtime contexts.
// Each method corresponds to a specific runtime (CLR, JVM, .NET Core, Perl, Ruby, Node.js, Python) and returns a RuntimeContext instance for that runtime.
type ConfigRuntimeFactory struct {
	path_ string
}

// NewConfigRuntimeFactory creates a new instance of RuntimeFactory with the specified configuration file path.
func NewConfigRuntimeFactory(path string) *ConfigRuntimeFactory {
	return &ConfigRuntimeFactory{path_: path}
}

// Clr creates a RuntimeContext instance to interact with CLR runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Clr(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Clr, configName)
}

// Jvm creates a RuntimeContext instance to interact with JVM runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Jvm(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Jvm, configName)
}

// Netcore creates a RuntimeContext instance to interact with .NET runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Netcore(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Netcore, configName)
}

// Perl creates a RuntimeContext instance to interact with Perl runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Perl(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Perl, configName)
}

// Python creates a RuntimeContext instance to interact with Python runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Python(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Python, configName)
}

// Ruby creates a RuntimeContext instance to interact with Ruby runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Ruby(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Ruby, configName)
}

// Nodejs creates a RuntimeContext instance to interact with Node.js runtime.
// configName - the name of the configuration to use.
// Refer to this article on Javonet Guides: https://www.javonet.com/guides/v2/golang/foundations/runtime-context
func (crf *ConfigRuntimeFactory) Nodejs(configName string) (*runtimecontext.RuntimeContext, error) {
	return crf.getRuntimeContext(runtimename.Nodejs, configName)
}

func (crf *ConfigRuntimeFactory) getRuntimeContext(runtimeName byte, configName string) (*runtimecontext.RuntimeContext, error) {
	if configName == "" {
		configName = "default"
	}

	jfr, err := jsonfileresolver.NewJsonFileResolver(crf.path_)
	if err != nil {
		return nil, err
	}

	licenseKey, err := jfr.GetLicenseKey()
	if err != nil {
		// licenseKey not found - do nothing
	} else {
		transmitter.ActivateWithCredentials(licenseKey)
	}
	runtimeNameStr, err := runtimenamehandler.GetName(runtimeName)
	if err != nil {
		return nil, err
	}
	connType, err := jfr.GetChannelType(runtimeNameStr, configName)
	if err != nil {
		return nil, err
	}

	if connType == "tcp" {
		runtimeNameStr, err := runtimenamehandler.GetName(runtimeName)
		if err != nil {
			return nil, err
		}
		host, err := jfr.GetChannelHost(runtimeNameStr, configName)
		if err != nil {
			return nil, err
		}

		port, err := jfr.GetChannelPort(runtimeNameStr, configName)
		if err != nil {
			return nil, err
		}

		connData := connectiondata.NewTcpConnectionData(host, (uint16(port)))
		return runtimecontext.GetInstance(runtimeName, connectiontype.Tcp, *connData)
	} else if connType == "inMemory" {
		return runtimecontext.GetInstance(runtimeName, connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0))
	} else {
		return nil, errors.New("Connection type " + connType + " not supported")
	}
}
