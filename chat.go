package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func ReadChat(conn net.Conn) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println(input.Text())
	}
}

func StartChatClient(Dial, ConnStr string) {
	conn, err := net.Dial(Dial, ConnStr)
	if err != nil {
		log.Fatal(err)
	}

	// 读取连接发送的消息
	go ReadChat(conn)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("you:")
		fmt.Fprintln(conn, input.Text())
	}
}
