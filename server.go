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
		buf := make([]byte, 512)
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

	result := query(buf)

	decode(result)

	pc.WriteTo(result, addr)
}

func query(request []byte) []byte {
	conn, err := net.Dial("udp", "223.5.5.5:53")
	defer conn.Close()
	if err != nil {
		logger.Panic(err)
	}
	conn.Write(request)
	response := make([]byte, 512)
	n, e := conn.Read(response)
	if e != nil {
		logger.Panic(e)
	}
	return response[:n]
}

func decode(buf []byte) {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	logger.Debugf("result: % x, length: %d", buf, len(buf))
	message := dns.Decode(buf)
	logger.Infow("", "message", message)
}
