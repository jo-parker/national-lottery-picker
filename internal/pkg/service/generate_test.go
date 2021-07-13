package service

import (
	"testing"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTicket(t *testing.T) {
	var mainCount, specialCount int
	type args struct {
		d model.Draw
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "EuroMillions OddEven",
			args: args{
				model.Draw{
					Name:       model.EuroMillions,
					NumTickets: 1,
					Day:        model.Tuesday,
					Strategy:   model.OddEven,
				},
			},
			wantErr: false,
		},
		{
			name: "Lotto OddEven",
			args: args{
				model.Draw{
					Name:       model.Lotto,
					NumTickets: 1,
					Day:        model.Tuesday,
					Strategy:   model.OddEven,
				},
			},
			wantErr: false,
		},
		{
			name: "Unknown Game",
			args: args{
				model.Draw{
					NumTickets: 1,
					Day:        model.Tuesday,
					Strategy:   model.OddEven,
				},
			},
			wantErr: true,
		},
		{
			name: "Unknown Strategy",
			args: args{
				model.Draw{
					Name:       model.Lotto,
					NumTickets: 1,
					Day:        model.Tuesday,
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateTicket(tt.args.d)

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			} else if (err != nil) == tt.wantErr {
				return
			}

			switch tt.args.d.Name {
			case model.EuroMillions:
				mainCount = 5
				specialCount = 2
			case model.Lotto:
				mainCount = 6
				specialCount = 0
			}

			medianBall := got.Game.NumMainBalls / 2
			evenCount, highCount := getBallNumberDistribution(got)

			assert.Equal(t, len(got.MainNumbers), mainCount)
			assert.Equal(t, len(got.SpecialNumbers), specialCount)
			assert.True(t, evenCount == medianBall || evenCount == medianBall+1)
			assert.True(t, highCount == medianBall || highCount == medianBall+1)
		})
	}
}

func getBallNumberDistribution(t *BaseTicket) (evenCount int, highCount int) {
	medianBallNumber := t.Game.MaxMainBall / 2

	var ec, hc int = 0, 0
	for number := range t.MainNumbers {
		if number%2 == 0 {
			ec++
		}

		if number > medianBallNumber {
			hc++
		}
	}

	return ec, hc
}
