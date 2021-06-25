package service

import (
	"testing"
	"os"

	"github.com/stretchr/testify/assert"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var euroMillionsDraw, lottoDraw model.Draw

func TestMain(m *testing.M) {
	euroMillionsDraw = model.Draw {
		Name: model.EuroMillions,
		NumTickets: 1,
		Day: model.Tuesday,
		Strategy: model.OddEven,
	}

	lottoDraw = model.Draw {
		Name: model.Lotto,
		NumTickets: 1,
		Day: model.Wednesday,
		Strategy: model.OddEven,
	}

	os.Exit(m.Run())
}

func TestGenerateTicketEuroMillions(t *testing.T) {
	ticket, _ := GenerateTicket(&euroMillionsDraw)

	medianBall := ticket.Game.NumMainBalls / 2
	evenCount, highCount := getBallNumberDistribution(ticket)

	assert.Equal(t, len(ticket.MainNumbers), 5)
	assert.Equal(t, len(ticket.SpecialNumbers), 2)

	assert.True(t, evenCount == medianBall || evenCount == medianBall + 1)
	assert.True(t, highCount == medianBall || highCount == medianBall + 1)
}

func TestGenerateTicketLotto(t *testing.T) {
	ticket, _ := GenerateTicket(&lottoDraw)

	medianBall := ticket.Game.NumMainBalls / 2
	evenCount, highCount := getBallNumberDistribution(ticket)

	assert.Equal(t, len(ticket.MainNumbers), 6)
	assert.Equal(t, len(ticket.SpecialNumbers), 0)

	assert.True(t, evenCount == medianBall || evenCount == medianBall + 1)
	assert.True(t, highCount == medianBall || highCount == medianBall + 1)
}

func getBallNumberDistribution(t *model.BaseTicket) (evenCount int, highCount int){
	medianBallNumber := t.Game.MaxMainBall / 2

	var ec, hc int = 0, 0
	for number, _ := range t.MainNumbers {
		if number % 2 == 0 {
			ec++
		}

		if number > medianBallNumber {
			hc++
		}
	}

	return ec, hc
}
