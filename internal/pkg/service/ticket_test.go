package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var (
	euroMillionsTicket, lottoTicket *BaseTicket = new(BaseTicket), new(BaseTicket)
)

func TestOddEvenTicket_SetBallNumbers(t *testing.T) {
	type fields struct {
		BaseTicket *BaseTicket
	}

	euroMillionsTicket.InitEuroMillionsTicket()
	lottoTicket.InitLottoTicket()

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "EuroMillions ticket",
			fields: fields{
				BaseTicket: euroMillionsTicket,
			},
		},
		{
			name: "Lotto ticket",
			fields: fields{
				BaseTicket: lottoTicket,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &OddEvenTicket{
				BaseTicket: tt.fields.BaseTicket,
			}
			tr.SetBallNumbers()

			switch tr.Game.Name {
			case model.EuroMillions:
				assert.Equal(t, len(tr.MainNumbers), 5)
				assert.Equal(t, len(tr.SpecialNumbers), 2)
			case model.Lotto:
				assert.Equal(t, len(tr.MainNumbers), 6)
				assert.Equal(t, len(tr.SpecialNumbers), 0)
			}
		})
	}
}
