package model

import ()

type Ticket struct {
	MainNumbers     map[int]struct{}
	SpecialNumbers  map[int]struct{}
	Entered         bool
	Draw            *Draw
}

func (t *Ticket) Init() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Entered = false
}

func (t *Ticket) InitEuromillions() {
	t.Draw = &Draw{
		Name: Euromillions,
		NumEvenBalls: 3,
		NumMainBalls: 5,
		NumSpecialBalls: 2,
	}
}

func (t *Ticket) InitLotto() {
	t.Draw = &Draw{
		Name: Lotto,
		NumEvenBalls: 3,
		NumMainBalls: 6,
		NumSpecialBalls: 0,
	}
}
