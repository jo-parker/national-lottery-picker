package model

import ()

type Day string

const (
	Monday Day = "mon"
	Tuesday Day = "tue"
	Wednesday Day = "wed"
	Thursday Day = "thu"
	Friday Day = "fri"
	Saturday Day = "sat"
	Sunday Day = "sun"
)

type Draw struct {
	Name                DrawName
	NumEvenBalls        int
	NumMainBalls        int
	NumSpecialBalls     int
	MaxSpecialBall			int
}

type DrawName string

const (
	EuroMillions DrawName = "euromillions"
	Lotto DrawName = "lotto"
)

var EuroMillionsDays = map[Day]struct{}{
	Tuesday: struct{}{},
	Friday: struct{}{},
}
var LottoDays = map[Day]struct{}{
	Wednesday: struct{}{},
	Saturday: struct{}{},
}
