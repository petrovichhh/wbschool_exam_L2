package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestCut(t *testing.T) {
	input := "field1\tfield2\tfield3\nfield4\tfield5\tfield6"
	expected := "field2\tfield3\nfield5\tfield6"

	r, w, _ := os.Pipe()
	oldStdout := os.Stdout
	os.Stdout = w

	cut(strings.NewReader(input), "\t", []int{2, 3}, false)

	w.Close()
	os.Stdout = oldStdout

	out, _ := ioutil.ReadAll(r)
	result := strings.TrimSuffix(string(out), "\n")
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
