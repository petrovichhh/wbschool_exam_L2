package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

type State interface {
	Handle() error
}

type WaitingState struct{}

func (s *WaitingState) Handle() error {
	fmt.Println("Waiting for input")
	return nil
}

type ServingState struct{}

func (d *ServingState) Handle() error {
	fmt.Println("Serving the drink")
	return nil
}

type VendingMachine struct {
	state State
}

func (v *VendingMachine) SetState(state State) {
	v.state = state
}

func (v *VendingMachine) Handle() error {
	return v.state.Handle()
}
