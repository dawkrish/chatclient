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

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("send: ")
		clientInp, err := clientReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		conn.Write([]byte(clientInp))

		serverResp, err := serverReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print("recv :", serverResp)
	}
}
