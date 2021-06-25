package service

import (
	"fmt"
	"errors"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

func GenerateTicket(d *model.Draw) (*model.BaseTicket, error) {
	base := new(model.BaseTicket)

	switch d.Name {
	case model.EuroMillions:
		base.InitEuroMillionsTicket()
	case model.Lotto:
		base.InitLottoTicket()
	default:
		return nil, errors.New(fmt.Sprintf("Unknown draw %v, exiting...", d.Name))
	}

	switch d.Strategy {
	case model.OddEven:
		t := new(model.OddEvenTicket)

		t.BaseTicket = base
		t.SetBallNumbers()
	case model.HotCold:
		t := new(model.HotColdTicket)

		t.BaseTicket = base
		t.SetBallNumbers()
	default:
		return nil, errors.New(fmt.Sprintf("Unknown strategy %v, exiting...", d.Strategy))
	}

	return base, nil
}
