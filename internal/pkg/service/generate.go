package service

import (
  "math/rand"
  "time"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var count, evenCount, oddCount int = 0, 0, 0

func init() {
	rand.Seed(time.Now().Unix())
}

func addMainBall(ballNumber int, t model.Ticket) *model.Ticket {
	medianBall := t.Game.NumMainBalls / 2 + 1

	evenAndBelowThreshold := (ballNumber % 2 == 0) && (evenCount < medianBall)
	oddAndBelowThreshold := (ballNumber % 2 != 0) && (oddCount < t.Game.NumMainBalls - medianBall)

	if evenAndBelowThreshold {
		t.MainNumbers[ballNumber] = struct{}{}

		if count < len(t.MainNumbers) {
			evenCount++
			count++
		}
	} else if oddAndBelowThreshold {
		t.MainNumbers[ballNumber] = struct{}{}

		if count < len(t.MainNumbers) {
			oddCount++
			count++
		}
	}

	return &t
}

func GenerateTicket() model.Ticket {
	count, evenCount, oddCount = 0, 0, 0

	var t *model.Ticket
	switch Config.NationalLottery.Game {
	case model.EuroMillions:
		t = model.NewEuroMillionsTicket()
	case model.Lotto:
		t = model.NewLottoTicket()
	}

	medianBallNumber := t.Game.MaxMainBall / 2
	for count < t.Game.NumMainBalls {
		var ballNumber int

		if count % 2 == 0 {
			ballNumber = 1 + rand.Intn(medianBallNumber)
		} else {
			ballNumber = 1 + medianBallNumber + rand.Intn(medianBallNumber)
		}

		addMainBall(ballNumber, *t)
	}

	for len(t.SpecialNumbers) < t.Game.NumSpecialBalls {
		ballNumber := 1 + rand.Intn(t.Game.MaxSpecialBall)
		t.SpecialNumbers[ballNumber] = struct{}{}
	}

	return *t
}
