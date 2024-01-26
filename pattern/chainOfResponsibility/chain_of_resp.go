package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	SetNext(Handler)
	Handle(string)
}

type ConcretedHandler1 struct {
	next Handler
}

func (h *ConcretedHandler1) SetNext(next Handler) {
	h.next = next
}

func (h *ConcretedHandler1) Handle(request string) {
	if request == "run" {
		fmt.Println("ConcretedHandler1: Running")

	} else if h.next != nil {
		h.next.Handle(request)
	}
}

type ConcretedHandler2 struct {
	next Handler
}

func (h *ConcretedHandler2) SetNext(next Handler) {
	h.next = next
}

func (h *ConcretedHandler2) Handle(request string) {
	if request == "stop" {
		fmt.Println("ConcretedHandler1: Stopping")

	} else if h.next != nil {
		h.next.Handle(request)
	}
}
