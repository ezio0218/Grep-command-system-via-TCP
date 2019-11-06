package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"strings"
)

type ClientRequest struct {
	Greppattern string
	LogFile     string
}

type Service string

func (serv *Service) Response(query ClientRequest, res *string) error {
	err := os.Chdir("/home/yaoxiao9/ourmp1/logfile/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	name := "grep"
	fmt.Println("Current Log File: ", query.LogFile)
	cmd := exec.Command(name, "-c", query.Greppattern, query.LogFile)
	fmt.Println(query.Greppattern)
	showCmd(cmd)
	out, error := cmd.Output()
	if error != nil {
		fmt.Println("Command Fails", error)
	}
	*res = strings.TrimSpace(string(out))
	return nil
}

func main() {
	service := new(Service)
	rpc.Register(service)
	address, error := net.ResolveTCPAddr("tcp", ":9000")
	if error != nil {
		log.Fatal("Address resolving Error: ", error)
		os.Exit(1)
	}
	listener, error := net.ListenTCP("tcp", address)
	fmt.Println("Server started")
	if error != nil {
		log.Fatal("Listening establishment Error: ", error)
		os.Exit(1)
	}
	for {
		conn, error := listener.Accept()
		if error != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
func showCmd(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}
