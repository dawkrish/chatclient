package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
)

var (
	clients = []net.Conn{}
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	networkType := os.Getenv("TYPE")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	ln, err := net.Listen(networkType, host+":"+port)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("Starting %v Server on %v:%v\n", networkType, host, port)

	for {
		conn, err := ln.Accept()
		fmt.Println("Incoming conn: ", conn.LocalAddr(), conn)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		clients = append(clients, conn)
		go reader(conn)
	}
}

func reader(conn net.Conn) {
	reader := bufio.NewReader(conn)
	closed := false
	for {
		requestFromClient, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection Closed:", requestFromClient)
			closed = true
		}
		for _, c := range clients {
			if c == conn {
				continue
			}
			n, err := c.Write([]byte(requestFromClient))
			if err != nil {
				panic(err)
			}
			fmt.Println("bytes written to ", c, " ", n)
		}
		if closed {
			break
		}
	}
}
