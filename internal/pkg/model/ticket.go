package model

type BaseTicket struct {
	MainNumbers    map[int]struct{}
	SpecialNumbers map[int]struct{}
	Game           *Game
}

type OddEvenTicket struct {
	*BaseTicket
}

type HotColdTicket struct {
	*BaseTicket
}

func (t *BaseTicket) InitEuroMillionsTicket() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Game = &Game{
		Name:            EuroMillions,
		NumMainBalls:    5,
		NumSpecialBalls: 2,
		MaxMainBall:     50,
		MaxSpecialBall:  12,
	}
}

func (t *BaseTicket) InitLottoTicket() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Game = &Game{
		Name:            Lotto,
		NumMainBalls:    6,
		NumSpecialBalls: 0,
		MaxMainBall:     59,
		MaxSpecialBall:  0,
	}
}
