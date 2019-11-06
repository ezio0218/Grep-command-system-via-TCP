package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
type ServerInfo struct {
	ServerId  int
	ServerAdr string
}
func main(){
	get_server_info("/Users/xy/go/src/MP1/ServerList/servers_adr.txt")
}
func get_server_info(servers_adr string) []ServerInfo {
	file_adr, err := os.Open(servers_adr)
	if err != nil {
		log.Fatal("error with reading servers address file!", err)
	}
	defer file_adr.Close()
	buf := bufio.NewReader(file_adr)
	server_id := 0
	servers_info := []ServerInfo{}
	for true {
		address, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal("error with reading servers address file!", err)
			}
		}
		server_id += 1
		//fmt.Print(address)
		address = strings.TrimSpace(address)
		server := ServerInfo{server_id, address}
		servers_info = append(servers_info, server)
	}
	fmt.Println(len(servers_info))
	return servers_info
}
