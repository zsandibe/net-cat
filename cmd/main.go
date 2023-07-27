package main

import (
	"fmt"
	"log"
	"net"
	"os"
	h "zsandibe/internal"
)

func main() {
	port := ":"
	args := os.Args[1:]

	if len(args) == 1 {
		port = port + args[0]
	} else if len(args) == 0 {
		port = ":8989"
	} else {
		log.Println("[USAGE]: ./TCPChat $port")
	}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Can`t listen on port: ", port)
	}
	fmt.Printf("Listening on the port %s\n", port)

	ch1 := make(chan h.Message)

	go h.Hub(ch1)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		h.Connections++
		if h.Connections <= 10 {
			go h.HandleConnection(conn, ch1)
		} else {
			conn.Write([]byte("Server is busy. Please try later.\n"))
			conn.Close()
		}
	}
}
