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

	t.Init()
	switch t.Draw.Name{
	case model.Euromillions:
		t.InitEuromillions()
	case model.Lotto:
		t.InitLotto()
	}

	for len(t.MainNumbers) < t.Draw.NumMainBalls {
		number := rand.Intn(10) + (count * 10)
		if number != 0 {
			addMainBall(number, *t)
		}
	}

	for len(t.SpecialNumbers) < t.Draw.NumSpecialBalls {
		specialBallNumber := 1 + rand.Intn(12)
		t.SpecialNumbers[specialBallNumber] = struct{}{}
	}

	return *t
}
