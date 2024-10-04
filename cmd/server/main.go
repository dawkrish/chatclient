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

	msgs := make(chan string)

	for {
		conn, err := ln.Accept()
		fmt.Println("Incoming conn: ", conn.LocalAddr())
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		clients = append(clients, conn)
		go reader(conn, msgs)

		go func() {
			for {
				msgRec := <-msgs
				if msgRec == "exit" {
					return
				}
				for _, c := range clients {
					if c == conn {
						continue
					}
					n, err := c.Write([]byte(msgRec))
					if err != nil {
						panic(err)
					}
					fmt.Println("bytes written to ", c, " ", n)
				}
			}
		}()
	}
}

func reader(conn net.Conn, msgs chan string) {
	for {
		reader := bufio.NewReader(conn)
		requestFromClient, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection Closed")
			msgs <- "exit"
			return
		}
		msgs <- requestFromClient
	}
}
