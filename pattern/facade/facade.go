package facade

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type CPU struct{}

func (c *CPU) ProcessesRequests() {
	fmt.Println("Query Processing..")
}

type HardDrive struct{}

func (h HardDrive) SaveInfo() {
	fmt.Println("Save info..")
}

type Memory struct{}

func (m Memory) RememberDevices() {
	fmt.Println("Remember devices..")
}

type Computer struct {
	cpu       *CPU
	hardDrive *HardDrive
	memory    *Memory
}

func (c *Computer) Start() {
	c.cpu.ProcessesRequests()
	c.hardDrive.SaveInfo()
	c.memory.RememberDevices()
}
