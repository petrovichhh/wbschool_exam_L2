package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

const (
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
}

type NoteBook struct {
	Type string
}

func NewNotebook() Computer {
	return &NoteBook{
		Type: NotebookType,
	}
}

func (n *NoteBook) GetType() string {
	return n.Type
}

type PersonalComputer struct {
	Type string
}

func NewPersonalComputer() Computer {
	return &PersonalComputer{
		Type: PersonalComputerType,
	}
}

func (c *PersonalComputer) GetType() string {
	return c.Type
}

func New(typeName string) Computer {
	switch typeName {
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	default:
		fmt.Printf("Non-existent object type %s", typeName)
		return nil
	}
}
