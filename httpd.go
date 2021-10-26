// Package main
package main

import (
	"net"
)

func main() {
	listener, err := StartUp(":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		conn.Write([]byte("xxx"))
		conn.Close()
	}

}

func AcceptRequest(client int) {

}

func StartUp(port string) (net.Listener, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		return nil, err
	}
	l, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		return nil, err
	}
	return l, nil
}
