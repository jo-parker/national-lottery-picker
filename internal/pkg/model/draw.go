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

type DrawName int

const (
	Euromillions DrawName = iota
	Lotto
)

func (d DrawName) String() string {
	return [...]string{"Euromillions", "Lotto"}[d]
}