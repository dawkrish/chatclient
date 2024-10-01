package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
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

	conn, err := ln.Accept()
	fmt.Println("Incoming conn: ", conn.LocalAddr())
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	reader(conn)
}

func reader(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		requestFromClient, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection Closed")
			return
		}

		n, err := conn.Write([]byte(requestFromClient))
		if err != nil {
			panic(err)
		}
		fmt.Println("bytes written :", n)
	}
}
