package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Data struct {
	dataFile   [][]string
	dataInput  [][]string
	dataSorted [][]string
}

func (d *Data) readData(file string) error {
	fileO, err := os.Open(file)
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}

	scanner := bufio.NewScanner(fileO)
	for scanner.Scan() {
		d.dataFile = append(d.dataFile, strings.Fields(scanner.Text()))
	}
	return nil
}

func (d *Data) readStdin() error {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		d.dataInput = append(d.dataInput, strings.Fields(input.Text()))
	}
	return input.Err()
}

func printData(data [][]string) {
	for _, line := range data {
		fmt.Println(strings.Join(line, " "))
	}
}

func handleErr(err error) {
	if err != nil {
		log.Println(nil)
	}
}

func columnToSort(data [][]string, column int, numeric, reverse bool) [][]string {
	sort.Slice(data, func(i, j int) bool {
		if column > 0 {
			column--
		}
		if len(data[i]) <= column {
			return true
		}
		if len(data[j]) <= column {
			return false
		}
		if numeric {
			in, _ := strconv.Atoi(data[i][column])
			jn, _ := strconv.Atoi(data[j][column])
			if reverse {
				return in > jn
			}
			return in < jn
		}
		if reverse {
			return data[i][column] > data[j][column]
		}
		return data[i][column] < data[j][column]
	})
	return data
}

func removeDuplicates(data [][]string) [][]string {
	result := [][]string{}
	seen := map[string]bool{}
	for _, line := range data {
		str := strings.Join(line, " ")
		if !seen[str] {
			result = append(result, line)
			seen[str] = true
		}
	}
	return result
}

func main() {
	flag_k := flag.Int("k", 1, "specifying the column to sort")
	flag_n := flag.Bool("n", false, "sort by numeric value")
	flag_r := flag.Bool("r", false, "sort in reverse order")
	flag_u := flag.Bool("u", false, "do not print duplicate lines")

	flag.Parse()

	d := Data{}

	if len(flag.Args()) > 0 {
		file := flag.Arg(0)
		err := d.readData(file)
		handleErr(err)
	} else {
		err := d.readStdin()
		handleErr(err)
	}

	if len(flag.Args()) > 0 {
		d.dataSorted = columnToSort(d.dataFile, *flag_k, *flag_n, *flag_r)
	} else {
		d.dataSorted = columnToSort(d.dataInput, *flag_k, *flag_n, *flag_r)
	}
	if *flag_u {
		d.dataSorted = removeDuplicates(d.dataSorted)
	}
	printData(d.dataSorted)
}
