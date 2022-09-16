package main

type Payment interface {
	Pay() error
}

type CardPayment struct {

}

func (p *CardPayment) {
	// implementation

}

type PayPalPayment struct {

}

func (p *PayPalPayment) {
	// implementation

}

type QiwiPayment struct {

}

func (p *QiwiPayment) {
	// implementation

}
