package main

import (
	"log"
	"github.com/liubaotong/socks5-proxy/socks5"
)

func main() {
	config := &socks5.Config{
		Username: "user",
		Password: "pass",
		Port:     1080,
	}

	server := socks5.NewServer(config)
	log.Fatal(server.Start())
}