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
	defer conn.Close()

	go serverReader(conn)
	clientReader(conn)
}

func clientReader(conn net.Conn) {
	clientReader := bufio.NewReader(os.Stdin)
	fmt.Print("name> ")
	name, err := clientReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = name[:len(name)-1]

	for {
		clientInp, err := clientReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if clientInp == "quit\n" || clientInp == "exit\n" {
			fmt.Printf("See you later\n")
			conn.Write([]byte(name + " has exited!"))
			conn.Close()
			break
		}
		sentence := name + ": " + clientInp
		conn.Write([]byte(sentence))
	}
}

func serverReader(conn net.Conn) {
	serverReader := bufio.NewReader(conn)
	for {
		serverResp, err := serverReader.ReadString('\n')
		if err != nil {
			conn.Close()
			return
		}
		fmt.Print(serverResp)
	}
}
