// Package main
package main

import (
	"net"
)

func main() {

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
