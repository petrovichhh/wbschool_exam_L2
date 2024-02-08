package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func readFile(file string) (string, error) {

	fileNew, err := os.Open(file)
	if err != nil {
		return "", errors.New("grepUt: text.txt: No such file or directory")
	}
	defer fileNew.Close()

	wr := bytes.Buffer{}
	sc := bufio.NewScanner(fileNew)

	for sc.Scan() {
		wr.WriteString(sc.Text())
		wr.WriteString("\n")
	}
	return wr.String(), nil
}

func printLines(w io.Writer, lines []string, args string, ignoreCase bool) {
	for _, line := range lines {
		if ignoreCase && strings.Contains(strings.ToLower(line), strings.ToLower(args)) {
			line = strings.ReplaceAll(line, args, "\033[31m"+args+"\033[0m")
			fmt.Fprint(w, line+"\n")
		} else if !ignoreCase && strings.Contains(line, args) {
			line = strings.ReplaceAll(line, args, "\033[31m"+args+"\033[0m")
			fmt.Fprint(w, line+"\n")
		}
	}
}

func grep() {
	flagI := flag.Bool("i", false, "ignore-case")
	flag.Parse()

	var args string
	var argFile string

	if !*flagI {
		args = os.Args[1]
		argFile = os.Args[2]
	} else {
		args = os.Args[2]
		argFile = os.Args[3]
	}

	data, err := readFile(argFile)
	if err != nil {
		fmt.Println(err)
	}
	lines := strings.Split(data, "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	printLines(os.Stdout, lines, args, *flagI)
}

func main() {
	grep()
}
