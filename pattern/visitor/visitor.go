package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Element struct {
	Accept (Checker)
}

type Checker interface {
	CheckCPU(*CPU)
	CheckHardDrive(*HardDrive)
	CheckMemory(*Memory)
}

type CPU struct{}

func (c *CPU) Accept(c Checker) {
	c.CheckCPU(c)
}

type HardDrive struct{}

func (h *HardDrive) Accept(c Checker) {
	c.CheckHardDrive(h)
}

type Memory struct{}

func (m *Memory) Accept(c Checker) {
	c.CheckMemory(m)
}
