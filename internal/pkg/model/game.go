package model

import ()

type Game struct {
	Name                GameName
	NumMainBalls        int
	NumSpecialBalls     int
	MaxMainBall					int
	MaxSpecialBall			int
}

type GameName string

const (
	EuroMillions GameName = "euromillions"
	Lotto GameName = "lotto"
)

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

var EuroMillionsDays = map[Day]struct{}{
	Tuesday: struct{}{},
	Friday: struct{}{},
}
var LottoDays = map[Day]struct{}{
	Wednesday: struct{}{},
	Saturday: struct{}{},
}
