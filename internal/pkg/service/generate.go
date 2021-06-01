package service

import (
  "math/rand"
  "time"
	"github.com/jpparker/euromillions-picker/internal/pkg/model"
)

const (
	midpoint = 25
	numLowBalls = 3
	totalMainBalls = 5
	totalSpecialBalls = 2
)

var count, evenCount, oddCount int = 0, 0, 0

func init() {
	rand.Seed(time.Now().Unix())
}

func addMainBall(ballNumber int, t model.Ticket) *model.Ticket {
	evenAndBelowThreshold := (ballNumber % 2 == 0) && (evenCount < 3)
	oddAndBelowThreshold := (ballNumber % 2 != 0) && (oddCount < 2)

	if evenAndBelowThreshold {
		t.MainNumbers[ballNumber] = struct{}{}

		if count != len(t.MainNumbers) {
			evenCount++
			count++
		}
	} else if oddAndBelowThreshold {
		t.MainNumbers[ballNumber] = struct{}{}

		if count != len(t.MainNumbers) {
			oddCount++
			count++
		}
	}

	return &t
}

func GenerateTicket() model.Ticket {
	count, evenCount, oddCount = 0, 0, 0

	ticket := new(model.Ticket)
	ticket.Init()

	for len(ticket.MainNumbers) != numLowBalls {
		lowBallNumber := 1 + rand.Intn(midpoint)
		addMainBall(lowBallNumber, *ticket)
	}

	for len(ticket.MainNumbers) != totalMainBalls {
		highBallNumber := midpoint + 1 + rand.Intn(midpoint)
		addMainBall(highBallNumber, *ticket)
	}

	for len(ticket.SpecialNumbers) != totalSpecialBalls {
		specialBallNumber := 1 + rand.Intn(12)
		ticket.SpecialNumbers[specialBallNumber] = struct{}{}
	}

	return *ticket
}
