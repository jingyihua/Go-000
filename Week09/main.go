package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3001")
	if err != nil {
		log.Fatalf("listen error:%v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error:%v\n", err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	rd := bufio.NewReader(conn)

	tmpChan := make(chan string, 8)
	go request(conn, tmpChan)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read error:%v\n", err)
			return
		}

		tmpChan <- string(line)
	}
}

func request(conn net.Conn, ch <-chan string) {
	wr := bufio.NewWriter(conn)

	for {
		select {
			case msg := <- ch:
				wr.WriteString(msg)
				wr.Flush()
		}
	}
}
