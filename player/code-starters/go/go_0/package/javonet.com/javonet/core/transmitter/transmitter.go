package transmitter

//Transmitter is used to transmit messages to native channel
type Transmitter struct {
}

func SendCommand(messageByteArray []byte) ([]byte, error) {
	return (&TransmitterWrapper{}).SendCommand(messageByteArray)
}

func ActivateWithLicenseFile() (int, error) {
	return activate("", "", "", "")
}

func ActivateWithCredentials(licenseKey string) (int, error) {
	return activate(licenseKey, "", "", "")
}

func ActivateWithCredentialsAndProxy(licenseKey string, proxyHost string, proxyUserName string, proxyPassword string) (int, error) {
	return activate(licenseKey, proxyHost, proxyUserName, proxyPassword)
}

func SetConfigSource(path string) (int, error) {
	return (&TransmitterWrapper{}).SetConfigSource(path)
}

func activate(licenseKey string, proxyHost string, proxyUserName string, proxyPassword string) (int, error) {
	return (&TransmitterWrapper{}).Activate(licenseKey, proxyHost, proxyUserName, proxyPassword)
}
