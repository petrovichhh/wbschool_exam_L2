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

/*
Применимость:
Когда нужно представить простой или урезанный интерфейс к сложной подсистеме.
Когда надо уменьшить количество зависимостей между клиентом и сложной системой.
Фасадные объекты позволяют отделить, изолировать компоненты системы от клиента и
развивать и работать с ними независимо.

Плюсы и минусы:
Осноным преимущество является изолируемость клиента от компонентов сложной подсистемы.
Недостатоком является риск стать выжным объектом, привязанным ко всем классам программы.

Реальный примеры использования:
Интернет магазин к которой подключена платежная система. Покупка товара.
*/
