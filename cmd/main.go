package main

import (
	"flag"
	"github.com/libuliduobuqiuqiu/chat-client"
)

var dial string
var connStr string

func init() {
	flag.StringVar(&dial, "d", "tcp", "For TCP, UDP and IP networks")
	flag.StringVar(&connStr, "c", "127.0.0.1:8080", "Connect IP address.")
}

func main() {
	flag.Parse()
	chat.StartChatClient(dial, connStr)
}
