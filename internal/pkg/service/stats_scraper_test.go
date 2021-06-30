package service

import (
	"testing"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestHotColdScraper(t *testing.T) {
	var mainCount, specialCount int

	type args struct {
		gn model.GameName
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "EuroMillions scraper",
			args: args{
				model.EuroMillions,
			},
			wantErr: false,
		},
		{
			name: "Lotto scraper",
			args: args{
				model.Lotto,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HotColdScraper(tt.args.gn)
			if (err != nil) != tt.wantErr {
				t.Errorf("HotColdScraper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			switch tt.args.gn {
			case model.EuroMillions:
				mainCount = 5
				specialCount = 2
			case model.Lotto:
				mainCount = 6
				specialCount = 0
			}

			assert.Equal(t, len(got[Hot].Main), mainCount)
			assert.Equal(t, len(got[Cold].Main), mainCount)
			assert.Equal(t, len(got[Hot].Special), specialCount)
			assert.Equal(t, len(got[Cold].Special), specialCount)
		})
	}
}
