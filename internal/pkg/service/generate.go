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
	evenAndBelowThreshold := (ballNumber % 2 == 0) && (evenCount < t.Draw.NumEvenBalls)
	oddAndBelowThreshold := (ballNumber % 2 != 0) && (oddCount < t.Draw.NumMainBalls - t.Draw.NumEvenBalls)

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

	t := new(model.Ticket)
	switch Config.NationalLottery.Draw {
	case model.EuroMillions:
		t.InitEuroMillions()
	case model.Lotto:
		t.InitLotto()
	}

	for len(t.MainNumbers) < t.Draw.NumMainBalls {
		ballNumber := rand.Intn(10) + (count * 10)
		if ballNumber != 0 {
			addMainBall(ballNumber, *t)
		}
	}

	for len(t.SpecialNumbers) < t.Draw.NumSpecialBalls {
		ballNumber := 1 + rand.Intn(t.Draw.MaxSpecialBall)
		t.SpecialNumbers[ballNumber] = struct{}{}
	}

	return *t
}
