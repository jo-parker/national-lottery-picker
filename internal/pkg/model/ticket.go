package model

import ()

type Ticket struct {
	MainNumbers				map[int]struct{}
	SpecialNumbers		map[int]struct{}
	Game							*Game
}

func NewEuroMillionsTicket() *Ticket {
	return &Ticket {
		MainNumbers: make(map[int]struct{}),
		SpecialNumbers: make(map[int]struct{}),
		Game: &Game{
			Name: EuroMillions,
			NumMainBalls: 5,
			NumSpecialBalls: 2,
			MaxMainBall: 50,
			MaxSpecialBall: 12,
		},
	}
}

func NewLottoTicket() *Ticket {
	return &Ticket {
		MainNumbers: make(map[int]struct{}),
		SpecialNumbers: make(map[int]struct{}),
		Game: &Game{
			Name: Lotto,
			NumMainBalls: 6,
			NumSpecialBalls: 0,
			MaxMainBall: 59,
			MaxSpecialBall: 0,
		},
	}
}
