package service

import (
	"fmt"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

func GenerateTicket(d model.Draw) (*BaseTicket, error) {
	base := new(BaseTicket)

	switch d.Name {
	case model.EuroMillions:
		base.InitEuroMillionsTicket()
	case model.Lotto:
		base.InitLottoTicket()
	default:
		return nil, fmt.Errorf("unknown draw %v, exiting", d.Name)
	}

	switch d.Strategy {
	case model.OddEven:
		t := new(OddEvenTicket)

		t.BaseTicket = base
		t.SetBallNumbers()
	case model.HotCold:
		t := new(HotColdTicket)

		t.BaseTicket = base
		t.SetBallNumbers()
	default:
		return nil, fmt.Errorf("unknown strategy %v, exiting", d.Strategy)
	}

	return base, nil
}
