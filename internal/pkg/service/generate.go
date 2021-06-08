package service

import (
  "math/rand"
  "time"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var count, evenCount, oddCount, pattern int = 0, 0, 0, 1

func init() {
	rand.Seed(time.Now().Unix())
}

func GenerateTicket(d *model.Draw, p int) *model.Ticket {
	count, evenCount, oddCount, pattern = 0, 0, 0, p

	var t *model.Ticket
	switch d.Name {
	case model.EuroMillions:
		t = model.NewEuroMillionsTicket()
	case model.Lotto:
		t = model.NewLottoTicket()
	}

	offset, remainder := patternVars()
	medianBallNumber := t.Game.MaxMainBall / 2

	for count < t.Game.NumMainBalls {
		var ballNumber int

		if count % 2 == remainder {
			ballNumber = 1 + rand.Intn(medianBallNumber)
		} else {
			ballNumber = 1 + medianBallNumber + rand.Intn(medianBallNumber)
		}

		addMainBall(ballNumber, t, offset)
	}

	for len(t.SpecialNumbers) < t.Game.NumSpecialBalls {
		ballNumber := 1 + rand.Intn(t.Game.MaxSpecialBall)
		t.SpecialNumbers[ballNumber] = struct{}{}
	}

	return t
}

func addMainBall(ballNumber int, t *model.Ticket, offset int) *model.Ticket {
	medianBall := t.Game.NumMainBalls / 2 + offset

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

	return t
}

func patternVars() (o int, r int) {
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
