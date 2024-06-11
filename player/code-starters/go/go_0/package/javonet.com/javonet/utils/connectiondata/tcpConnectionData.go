package connectiondata

import "net"

type TcpConnectionData struct {
	IpAddress net.IPAddr
	Port      uint16
}

func NewTcpConnectionData(hostname string, port uint16) *TcpConnectionData {
	ipAddress, err := net.ResolveIPAddr("ip", hostname)
	if err != nil {
		panic(err)
	}
	return &TcpConnectionData{IpAddress: *ipAddress, Port: port}
}

func (t *TcpConnectionData) GetAddressBytes() []byte {
	return t.IpAddress.IP.To4()
}

func (t *TcpConnectionData) GetPortBytes() []byte {
	return []byte{byte(t.Port), byte(t.Port >> 8)}
}
