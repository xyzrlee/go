package main

import (
	"flag"
	"net"

	"./dns"
	"github.com/xyzrlee/go/logger"
)

func init() {
	logger.SetLevel(logger.DebugLevel)
}

func main() {

	var a []string
	a = append(a, "asdf")
	logger.Debugw("", "a", a)

	var listen string
	flag.StringVar(&listen, "listen", ":53", "listen address")
	flag.Parse()

	logger.Infow("params", "listen", listen)

	pc, err := net.ListenPacket("udp", listen)
	if err != nil {
		logger.Fatal(err)
	}
	defer pc.Close()

	for {
		buf := make([]byte, 10240)
		n, addr, err := pc.ReadFrom(buf)
		if err != nil {
			logger.Error(err)
			continue
		}
		go Serve(pc, addr, buf[:n])
	}

}

func Serve(pc net.PacketConn, addr net.Addr, buf []byte) {
	defer logger.Sync()
	logger.Debugf("buf: % x, length: %d", buf, len(buf))
	message := dns.Decode(buf)
	logger.Infow("", "message", message)
	// pc.WriteTo(buf, addr)
}
