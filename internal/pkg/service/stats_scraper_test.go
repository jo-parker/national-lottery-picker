package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

func TestHotColdScraperLotto(t *testing.T) {
	numbers, _ := HotColdScraper(model.Lotto)

	assert.Equal(t, len(numbers[Hot].Main), 6)
	assert.Equal(t, len(numbers[Cold].Main), 6)

	assert.Equal(t, len(numbers[Hot].Special), 0)
	assert.Equal(t, len(numbers[Cold].Special), 0)
}

func TestHotColdScraperEuroMillions(t *testing.T) {
	numbers, _ := HotColdScraper(model.EuroMillions)

	assert.Equal(t, len(numbers[Hot].Main), 5)
	assert.Equal(t, len(numbers[Cold].Main), 5)

	assert.Equal(t, len(numbers[Hot].Special), 2)
	assert.Equal(t, len(numbers[Cold].Special), 2)
}
