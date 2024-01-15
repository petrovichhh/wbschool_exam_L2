package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

type Time interface {
	getTime() (time.Time, error)
}

type ntpTime struct{}

func (n ntpTime) getTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}

func main() {
	t := ntpTime{}
	tim, err := t.getTime()

	if err != nil {
		fmt.Fprintln(os.Stderr, "error getting time", err)
		os.Exit(1)
	}
	fmt.Println(tim)

	fmt.Println(time.Now())
}
