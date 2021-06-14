package model

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type BaseTicket struct {
	MainNumbers				map[int]struct{}
	SpecialNumbers		map[int]struct{}
	Game							*Game
}


func (t *BaseTicket) InitEuroMillionsTicket() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Game = &Game{
		Name: EuroMillions,
		NumMainBalls: 5,
		NumSpecialBalls: 2,
		MaxMainBall: 50,
		MaxSpecialBall: 12,
	}
}

func (t *BaseTicket) InitLottoTicket() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Game = &Game{
		Name: Lotto,
		NumMainBalls: 6,
		NumSpecialBalls: 0,
		MaxMainBall: 59,
		MaxSpecialBall: 0,
	}
}


type Ticket interface {
	SetBallNumbers()
	patternVars()		(int, int)
}

type OddEvenTicket struct {
	*BaseTicket
}

type HotColdTicket struct {
	*BaseTicket
}

func (t *OddEvenTicket) SetBallNumbers() {
	var totalCount, evenCount, oddCount int = 0, 0, 0

	offset, remainder := t.patternVars()
	medianBall := t.Game.NumMainBalls / 2 + offset
	medianBallNumber := t.Game.MaxMainBall / 2

	for totalCount < t.Game.NumMainBalls {
		var ballNumber int

		if totalCount % 2 == remainder {
			ballNumber = 1 + rand.Intn(medianBallNumber)
		} else {
			ballNumber = 1 + medianBallNumber + rand.Intn(medianBallNumber)
		}

		evenAndBelowThreshold := (ballNumber % 2 == 0) && (evenCount < medianBall)
		oddAndBelowThreshold := (ballNumber % 2 != 0) && (oddCount < t.Game.NumMainBalls - medianBall)

		if evenAndBelowThreshold {
			t.MainNumbers[ballNumber] = struct{}{}

			if totalCount < len(t.MainNumbers) {
				evenCount++
				totalCount++
			}
		} else if oddAndBelowThreshold {
			t.MainNumbers[ballNumber] = struct{}{}

			if totalCount < len(t.MainNumbers) {
				oddCount++
				totalCount++
			}
		}
	}

	for len(t.SpecialNumbers) < t.Game.NumSpecialBalls {
		ballNumber := 1 + rand.Intn(t.Game.MaxSpecialBall)
		t.SpecialNumbers[ballNumber] = struct{}{}
	}
}

func (t *OddEvenTicket) patternVars() (o int, r int) {
	pattern := rand.Intn(4)

	var offset, remainder int
	switch pattern {
	case 0:
		offset = 1; remainder = 0
	case 1:
		offset = 0; remainder = 0
	case 2:
		offset = 0; remainder = 1
	case 3:
		offset = 1; remainder = 1
	}

	return offset, remainder
}

func (t *HotColdTicket) SetBallNumbers() {}
