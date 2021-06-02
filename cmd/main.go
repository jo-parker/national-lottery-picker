package main

import (
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
)

func main() {
	draw := model.Euromillions
	service.EnterDraw(draw)
}
