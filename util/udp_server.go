package util

import (
	"github.com/sirupsen/logrus"
	"net"
)

type UdpServer struct {
	port    int
	conn    *net.UDPConn
	handler func(data []byte)
}

func (s *UdpServer) Serve() (err error) {

	for {
		var data [1024 * 1024]byte // 1MB
		n, addr, err := s.conn.ReadFromUDP(data[:])
		if err != nil {
			logrus.Errorf("recv udp err: %s", err.Error())
			return err
		}

		logrus.Infof("recv udp from addr: %s, size: %d", addr.String(), n)

		// TODO: 后续可以改造程多线程
		s.handler(data[:n])
	}
}

func (s *UdpServer) Shutdown() {
	_ = s.conn.Close()
	return
}

func (s *UdpServer) do(data []byte) {
	defer func() {
		if e := recover(); e != nil {
			logrus.Errorf("catch panic e: %v", e)
		}
	}()

	s.handler(data)
	return
}

func NewUdpServer(port int, handler func(data []byte)) *UdpServer {
	addr := &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: port,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logrus.Errorf("listen port=%d err: %s", port, err.Error())
		return nil
	}

	return &UdpServer{
		port:    port,
		conn:    conn,
		handler: handler,
	}
}
