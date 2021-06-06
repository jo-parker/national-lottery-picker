package model

import ()

type Ticket struct {
	MainNumbers     map[int]struct{}
	SpecialNumbers  map[int]struct{}
	Entered         bool
	Draw            *Draw
}

func (t *Ticket) InitEuroMillions() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Entered = false

	t.Draw = &Draw{
		Name: EuroMillions,
		NumEvenBalls: 3,
		NumMainBalls: 5,
		NumSpecialBalls: 2,
		MaxSpecialBall: 12,
	}
}

func (t *Ticket) InitLotto() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Entered = false

	t.Draw = &Draw{
		Name: Lotto,
		NumEvenBalls: 3,
		NumMainBalls: 6,
		NumSpecialBalls: 0,
		MaxSpecialBall: 0,
	}
}
