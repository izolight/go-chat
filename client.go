package main

import (
	"net"
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error:%s", err.Error())
		os.Exit(1)
	}
}

func main() {
	conn, err := net.Dial("udp", "localhost:9876")
	checkError(err)
	defer conn.Close()

}

