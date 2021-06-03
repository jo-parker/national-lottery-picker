package model

import ()

type Draw struct {
	Name                DrawName
	Midpoint            int
	NumLowBalls         int
	NumEvenBalls        int
	NumMainBalls        int
	NumSpecialBalls     int
}

type DrawName string

const (
	Euromillions DrawName = "euromillions"
	Lotto DrawName = "lotto"
)
