package service

import (
	"testing"
	"fmt"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

func TestMain2(t *testing.T) {
	numbers, _ := HotColdScraper(model.Lotto)

	fmt.Println((*numbers)[Hot])
	fmt.Println((*numbers)[Cold])
}
