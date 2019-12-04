package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "12345"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", HOST+":"+PORT)
	exitOnError(err)

	listener, err := net.ListenTCP("tcp", addr)
	exitOnError(err)

	fmt.Println("Listening port" + PORT + "...")

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			go handleClient(client)
		}
	}

}

func handleClient(client net.Conn) {
	defer fmt.Println("close client")
	defer client.Close()

	fmt.Println("Connected from:", client.RemoteAddr().String())

	for {
		buf := make([]byte, 512)
		_, err := client.Read(buf)
		if err == io.EOF {
			fmt.Println("recv close from client")
			return
		}

		if err != nil {
			fmt.Println("read error:", err)
			return
		}

		fmt.Println(string(buf))

		_, err = client.Write(buf)
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}
