package dns

import (
	"net"
)

type A struct {
	IP net.IP
}

func (a *A) Read(reader *reader, length uint16) ResourceData {
	a.IP = reader.read(int(length))
	return a
}

func (a *A) String() string {
	return a.IP.String()
}
