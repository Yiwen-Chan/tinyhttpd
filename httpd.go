// Package main
package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := StartUp("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		AcceptRequest(conn)
		conn.Close()
	}

}

func AcceptRequest(client net.Conn) error {
	var method string
	br := bufio.NewReader(client)
	b, _, err := br.ReadLine()
	if err != nil {
		return err
	}
	buf := string(b)
	i := strings.Index(buf, " ")
	if i == -1 {
		return errors.New("Unknown method")
	}
	method = buf[:i]
	fmt.Printf("v%sv", method)
	return nil
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
