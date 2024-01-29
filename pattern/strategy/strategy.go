package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type MoveStrategy interface {
	Move()
}

type Run struct{}

func (Run) Move() {
	fmt.Println("I run fast")
}

type Crawl struct{}

func (Crawl) Move() {
	fmt.Println("I crawl slowly")
}

type Character struct {
	moveStrategy MoveStrategy
}

func (c *Character) setMoveStrategy(m MoveStrategy) {
	c.moveStrategy = m
}

func (c *Character) Move() {
	c.moveStrategy.Move()
}
