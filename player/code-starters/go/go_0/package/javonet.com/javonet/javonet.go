package javonet

import (
	"javonet.com/javonet/core/transmitter"
	"javonet.com/javonet/sdk/configruntimefactory"
	"javonet.com/javonet/sdk/runtimefactory"
	"javonet.com/javonet/utils/connectiondata"
	"javonet.com/javonet/utils/connectiontype"
)

// Javonet is a singleton struct that serves as the entry point for interacting with Javonet.
// It provides methods to activate and initialize the Javonet SDK.
// It supports both in-memory and TCP connections.
// See also: https://www.javonet.com/guides/v2/golang/foundations/javonet-static-class
type Javonet struct {
}

func init() {
	activate()
}

// activate activates Javonet with the license file.
// Returns the activation status code and an error if any.
func activate() (int, error) {
	return transmitter.ActivateWithLicenseFile()
}

// ActivateWithCredentials activates Javonet with the provided license key.
// Returns the activation status code and an error if any.
// See also: https://www.javonet.com/guides/v2/golang/getting-started/activating-javonet
func ActivateWithCredentials(licenseKey string) (int, error) {
	return transmitter.ActivateWithCredentials(licenseKey)
}

// ActivateWithCredentrialsAndProxy activates Javonet with the provided license key and proxy host.
// Returns the activation status code and an error if any.
// See also: https://www.javonet.com/guides/v2/golang/getting-started/activating-javonet
func ActivateWithCredentialsAndProxy(licenseKey string, proxyHost string, proxyUserName string, proxyPassword string) (int, error) {
	return transmitter.ActivateWithCredentialsAndProxy(licenseKey, proxyHost, proxyUserName, proxyPassword)
}

// InMemory initializes Javonet using an in-memory channel on the same machine.
// Returns a RuntimeFactory instance configured for an in-memory connection.
// See also: https://www.javonet.com/guides/v2/golang/foundations/in-memory-channel
func InMemory() *runtimefactory.RuntimeFactory {
	return runtimefactory.NewRuntimeFactory(connectiontype.InMemory, *connectiondata.NewTcpConnectionData("0.0.0.0", 0))
}

// Tcp initializes Javonet with a TCP connection to a remote machine.
// Returns a RuntimeFactory instance configured for a TCP connection.
// See also: https://www.javonet.com/guides/v2/golang/foundations/tcp-channel
func Tcp(connectionData connectiondata.TcpConnectionData) *runtimefactory.RuntimeFactory {
	return runtimefactory.NewRuntimeFactory(connectiontype.Tcp, connectionData)
}

// WithConfig initializes Javonet with a configuration file.
// Currently supported: Configuration file in JSON format.
// Returns a ConfigRuntimeFactory instance with configuration data.
// See also: https://www.javonet.com/guides/v2/golang/foundations/configure-channel
func WithConfig(path string) *configruntimefactory.ConfigRuntimeFactory {
	_, err := transmitter.SetConfigSource(path)
	if err != nil {
		panic(err)
	}
	return configruntimefactory.NewConfigRuntimeFactory(path)
}
