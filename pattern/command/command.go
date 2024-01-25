package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Command interface {
	execute()
}

type Device struct {
	status bool
}

func (d *Device) turnOn() {
	d.status = true
	fmt.Println("Device is turned on")
}

func (d *Device) turnOff() {
	d.status = false
	fmt.Println("Device is turned off")
}

type TurnOnCommand struct {
	device *Device
}

func (c *TurnOnCommand) execute() {
	c.device.turnOn()
}

type TurnOffCommand struct {
	device *Device
}

func (c *TurnOffCommand) execute() {
	c.device.turnOff()
}

type RemoteControl struct {
	command Command
}

func (r *RemoteControl) setCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) pressButtom() {
	r.command.execute()
}
