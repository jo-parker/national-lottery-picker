package model

import ()

type Draw struct {
	Name                DrawName
	NumEvenBalls        int
	NumMainBalls        int
	NumSpecialBalls     int
}

type DrawName string

const (
	Euromillions DrawName = "euromillions"
	Lotto DrawName = "lotto"
)
