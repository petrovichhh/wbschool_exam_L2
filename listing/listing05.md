Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
Вывод: error
В Go, nil для конкретного типа и nil для интерфейса - это не одно и то же.  err объявлен как интерфейс error. Когда  возвращается nil в функции test(), он имеет тип *customError. Поэтому, когда проверяется err != nil, это условие оказывается истинным, потому что err не равен nil с точки зрения интерфейса error. 