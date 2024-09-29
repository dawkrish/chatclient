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

	conn, err := net.Dial(networkType, host+":"+port)
	if err != nil {
		panic(err)
	}
	for {
		fmt.Print("send: ")
		reader := bufio.NewReader(conn)
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(line))
	}
}
