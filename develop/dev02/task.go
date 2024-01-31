package main

import (
	"bytes"
	"errors"
	"fmt"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpackingString(str string) (string, error) {
	var buf bytes.Buffer
	stringRune := []rune(str)

	for i := 0; i < len(stringRune); i++ {

		if unicode.IsDigit(stringRune[i]) {
			if i == 0 {
				return "", errors.New("invalid line")
			}
			count := int(stringRune[i] - '0')
			for j := 0; j < count-1; j++ {
				buf.WriteRune(stringRune[i-1])
			}
		} else {
			buf.WriteRune(stringRune[i])
		}
	}
	return buf.String(), nil
}

func main() {
	arrString := []string{"a4bc2d5e", "abcd", "45", ""}
	for _, input := range arrString {
		output, err := unpackingString(input)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(output)
	}
}
