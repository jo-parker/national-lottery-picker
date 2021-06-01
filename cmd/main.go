package main

import (
	"github.com/jpparker/euromillions-picker/internal/pkg/model"
	"github.com/jpparker/euromillions-picker/internal/pkg/service"
)

func main() {
	draw := model.Euromillions
	service.EnterDraw(draw)
}
