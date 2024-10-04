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

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)
	fmt.Print("name> ")
	name, err := clientReader.ReadString('\n')
	name = name[:len(name)-1]
	if err != nil {
		panic(err)
	}
	//srvMsgs := make(chan string)
	go func() {
		for {
			serverResp, err := serverReader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			fmt.Print(serverResp)
		}
	}()
	for {
		clientInp, err := clientReader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if clientInp == "quit\n" || clientInp == "exit\n" {
			fmt.Printf("See you later\n")
			conn.Close()
			return
		}
		sentence := name + ": " + clientInp
		conn.Write([]byte(sentence))
	}
}
