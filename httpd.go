// Package main
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const SERVER_STRING = "Server: jdbhttpd/0.1.0\r\n"

func main() {
	server, err := StartUp("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		client, err := server.Accept()
		if err != nil {
			continue
		}
		AcceptRequest(client)
		client.Close()
	}
}

func AcceptRequest(client net.Conn) error {
	var (
		method string
		url    string
	)
	br := bufio.NewReader(client)
	method, err := br.ReadString(' ')
	if err != nil {
		return err
	}
	if method != "GET" && method != "POST" {
		UnimpleMented(client)
		return nil
	}

	i = strings.Index(buf[i+1:], " ")
	if i != -1 {
		url = buf[:i]
	} else {
		UnimpleMented(client)
		return nil
	}
	url = buf[i+1:]
	if method == "GET" {

	}
	fmt.Println(buf)
	return nil
}

func UnimpleMented(client net.Conn) {
	bw := bufio.NewWriterSize(client, 1024)
	bw.WriteString("HTTP/1.0 501 Method Not Implemented\r\n")
	bw.WriteString(SERVER_STRING)
	bw.WriteString("Content-Type: text/html\r\n")
	bw.WriteString("\r\n")
	bw.WriteString("<HTML><HEAD><TITLE>Method Not Implemented\r\n")
	bw.WriteString("</TITLE></HEAD>\r\n")
	bw.WriteString("<BODY><P>HTTP request method not supported.\r\n")
	bw.WriteString("</BODY></HTML>\r\n")
	bw.Flush()
}

func NotFound(client net.Conn) {
	bw := bufio.NewWriterSize(client, 1024)
	bw.WriteString("HTTP/1.0 404 NOT FOUND\r\n")
	bw.WriteString(SERVER_STRING)
	bw.WriteString("Content-Type: text/html\r\n")
	bw.WriteString("\r\n")
	bw.WriteString("<HTML><TITLE>Not Found</TITLE>\r\n")
	bw.WriteString("<BODY><P>The server could not fulfill\r\n")
	bw.WriteString("your request because the resource specified\r\n")
	bw.WriteString("is unavailable or nonexistent.\r\n")
	bw.WriteString("</BODY></HTML>\r\n")
	bw.Flush()
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
