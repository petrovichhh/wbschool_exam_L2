package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func removeDublicate(arr []string) []string {
	sort.Strings(arr)
	j := 0
	for i := 1; i < len(arr); i++ {
		if arr[j] != arr[i] {
			j++
			arr[j] = arr[i]
		}
	}
	return arr[:j+1]
}

func anagram(array []string) (map[string][]string, error) {
	data := make(map[string][]string)

	for _, str := range array {
		str = strings.ToLower(str)
		valueRune := []rune(str)

		sort.Slice(valueRune, func(i, j int) bool {
			return valueRune[i] < valueRune[j]
		})

		sortedStr := string(valueRune)

		for _, sec := range array {
			sec = strings.ToLower(sec)
			valueComparison := []rune(sec)

			sort.Slice(valueComparison, func(i, j int) bool {
				return valueComparison[i] < valueComparison[j]
			})

			if sortedStr == string(valueComparison) {
				data[sortedStr] = append(data[sortedStr], sec)
			}
		}
	}

	for key, value := range data {
		if len(value) < 2 {
			delete(data, key)
		} else {
			data[key] = removeDublicate(value)
		}
	}

	return data, nil
}

func main() {

	array := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столик"}
	data, err := anagram(array)
	if err != nil {
		fmt.Println(err)
	}
	for key, values := range data {
		fmt.Printf("key = %v values: \n", key)
		for _, value := range values {
			fmt.Println(value)
		}
	}
}
