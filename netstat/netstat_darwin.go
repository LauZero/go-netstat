//go:build darwin
// +build darwin

// TODO:
package netstat

var skStates = [...]string{}

// Socket states
const (
	Established SkState = 0x01
	SynSent             = 0x02
	SynRecv             = 0x03
	FinWait1            = 0x04
	FinWait2            = 0x05
	TimeWait            = 0x06
	Close               = 0x07
	CloseWait           = 0x08
	LastAck             = 0x09
	Listen              = 0x0a
	Closing             = 0x0b
)

// TCPSocks returns a slice of active TCP sockets containing only those
// elements that satisfy the accept function
func osTCPSocks(accept AcceptFn) ([]SockTabEntry, error) {
	return nil, nil
}

// TCP6Socks returns a slice of active TCP IPv4 sockets containing only those
// elements that satisfy the accept function
func osTCP6Socks(accept AcceptFn) ([]SockTabEntry, error) {
	return nil, nil
}

// UDPSocks returns a slice of active UDP sockets containing only those
// elements that satisfy the accept function
func osUDPSocks(accept AcceptFn) ([]SockTabEntry, error) {
	return nil, nil
}

// UDP6Socks returns a slice of active UDP IPv6 sockets containing only those
// elements that satisfy the accept function
func osUDP6Socks(accept AcceptFn) ([]SockTabEntry, error) {
	return nil, nil
}
