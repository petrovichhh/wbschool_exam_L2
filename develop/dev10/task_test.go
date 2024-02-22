package main

import (
	"net"
	"testing"
	"time"
)

func TestTelnetClient(t *testing.T) {
	// Запускаем тестовый сервер
	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		t.Fatalf("Cannot start test server: %v", err)
	}
	defer server.Close()

	go func() {
		conn, _ := server.Accept()
		defer conn.Close()
		conn.Write([]byte("Hello, client!"))
	}()

	// Запускаем тестовый клиент
	timeout := 1 * time.Second
	host := "localhost"
	port := "9999"

	dialer := &net.Dialer{
		Timeout: timeout,
	}

	conn, err := dialer.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		t.Fatalf("Cannot connect: %v", err)
	}
	defer conn.Close()

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		t.Fatalf("Error occurred while receiving data: %v", err)
	}

	expected := "Hello, client!"
	if string(buf[:n]) != expected {
		t.Fatalf("Expected '%s', got '%s'", expected, string(buf[:n]))
	}
}
