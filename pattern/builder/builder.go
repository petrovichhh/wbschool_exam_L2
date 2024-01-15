package builder

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Collector interface {
	SetCore(value int)
	SetMemory(value int)
	SetMonitor(value int)
	SetGraphicCard(value int)
}

type Fabric struct {
	collector Collector
}

func (f *Fabric) Construct() {
	f.collector.SetCore(6)
	f.collector.SetMemory(16)
	f.collector.SetMonitor(1)
	f.collector.SetGraphicCard(2)
}

type ConcreteBuilder struct {
	computer *Computer
}

func (c *ConcreteBuilder) SetCore(value int) {
	c.computer.Core = value
}
func (c *ConcreteBuilder) SetMemory(value int) {
	c.computer.Memory = value
}
func (c *ConcreteBuilder) SetMonitor(value int) {
	c.computer.Monitor = value
}
func (c *ConcreteBuilder) SetGraphicCard(value int) {
	c.computer.GraphicCard = value
}

type Computer struct {
	Core        int
	Memory      int
	Monitor     int
	GraphicCard int
}

func (pc *Computer) Show() {
	fmt.Printf(" Core:[%d] Mem:[%d] GraphicCard:[%d] Monitor:[%d]",
		pc.Core,
		pc.Memory,
		pc.GraphicCard,
		pc.Monitor)
}
