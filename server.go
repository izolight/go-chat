package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

type Server struct {
	conn *net.UDPConn
	//buf chan string
	clients [] *net.UDPAddr
}

func (s *Server)handleConnection() {
	buf := make([]byte, 1024)
	n, addr, err := s.conn.ReadFromUDP(buf)
	checkError(err)

	if !addrInSlice(addr, s.clients) {
		s.clients = append(s.clients, addr)
	}

	msg := addr.String() + ": " + string(buf[0:n])
	for _, a := range s.clients {
		if addr.String() != a.String() {
			_, err = s.conn.WriteToUDP([]byte(msg), a)
		}

		//checkError(err)
	}
}

func addrInSlice(a *net.UDPAddr, list [] *net.UDPAddr) bool {
	for _, b := range list {
		if a.String() == b.String() {
			return true
		}
	}
	return false
}

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":9876")
	checkError(err)

	var s Server
	s.conn, err = net.ListenUDP("udp", serverAddr)
	//s.buf = make(chan string, 10)
	s.clients = make([] *net.UDPAddr, 10)
	checkError(err)
	defer s.conn.Close()

	for {
		s.handleConnection()
	}
}
