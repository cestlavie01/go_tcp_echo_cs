package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

const PORT = 12345

func main() {
	portStr := strconv.Itoa(PORT)
	server, err := net.Listen("tcp", "0.0.0.0:"+portStr)
	if server == nil {
		panic("failed to listen:" + err.Error())
	}

	fmt.Println("Listening port:", portStr)

	clients := acceptClient(server)
	for {
		go handleClient(<-clients)
	}
}

func acceptClient(listner net.Listener) chan net.Conn {
	ch := make(chan net.Conn)

	go func() {
		for {
			client, err := listner.Accept()
			if client == nil {
				fmt.Println("failed to accept:", err.Error())
			}

			fmt.Printf("new connection: %v <-> %v\n", client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()

	return ch
}

func handleClient(client net.Conn) {
	defer fmt.Println("close client")
	defer client.Close()

	fmt.Println("Connected from:", client.RemoteAddr().String())

	reader := bufio.NewReader(client)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("failed to read from client.", err.Error())
			break
		}

		fmt.Println(string(line))

		_, err = client.Write(line)
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
	}
}
