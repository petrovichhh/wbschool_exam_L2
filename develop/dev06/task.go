package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func cut(input io.Reader, delimiter string, fields []int, separated bool) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()

		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		parts := strings.Split(line, delimiter)
		var output []string
		for _, field := range fields {
			if field > 0 && field <= len(parts) {
				output = append(output, parts[field-1])
			}
		}
		fmt.Println(strings.Join(output, "\t")) // выводим результат с TAB-разделителем
	}
}

func main() {
	var (
		fieldsFlag    string
		delimiterFlag string
		separatedFlag bool
	)
	flag.StringVar(&fieldsFlag, "f", "1", "Выбрать поля (колонки)")
	flag.StringVar(&delimiterFlag, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&separatedFlag, "s", false, "Только строки с разделителем")
	flag.Parse()

	fields := parseFields(fieldsFlag)

	cut(os.Stdin, delimiterFlag, fields, separatedFlag)
}

func parseFields(fieldsFlag string) []int {
	fieldsStr := strings.Split(fieldsFlag, ",")
	var fields []int
	for _, fieldStr := range fieldsStr {
		field := 0
		fmt.Sscanf(fieldStr, "%d", &field)
		fields = append(fields, field)
	}
	return fields
}
