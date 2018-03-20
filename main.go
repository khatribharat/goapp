package main

import (
	"net"
	"fmt"
	"os"
	"strconv"
)

// IsPortAvailable checks if a TCP port is available or not
func IsPortAvailable(p int) bool {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", p))
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func main() {
	arg := os.Args[1]
	port, err := strconv.Atoi(arg)
	if (err != nil) {
		fmt.Printf("The first and only argument should be a valid port: %s is invalid", arg)
	} else {
		fmt.Printf("Is port %s available? %t\n", arg, IsPortAvailable(port))
	}
}
