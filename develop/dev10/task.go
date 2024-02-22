package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		log.Fatalf("Usage: go-telnet [--timeout] host port")
	}

	host := args[0]
	port := args[1]

	dialer := &net.Dialer{
		Timeout: *timeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}
	defer conn.Close()

	go func() {
		if _, err := io.Copy(conn, os.Stdin); err != nil {
			log.Fatalf("Error occurred while sending data: %v", err)
		}
	}()

	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatalf("Error occurred while receiving data: %v", err)
	}
}
