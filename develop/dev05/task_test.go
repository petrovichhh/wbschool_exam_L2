package main

import (
	"bytes"
	"testing"
)

func TestPrintLines(t *testing.T) {
	lines := []string{"Hello", "World", "Hello", "Gopher"}
	args := "Hello"
	ignoreCase := false

	buf := new(bytes.Buffer)
	printLines(buf, lines, args, ignoreCase)

	want := "\033[31mHello\033[0m\n\033[31mHello\033[0m\n"
	got := buf.String()

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
