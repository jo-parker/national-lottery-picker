package model

type Game struct {
	Name            GameName
	NumMainBalls    int
	NumSpecialBalls int
	MaxMainBall     int
	MaxSpecialBall  int
}

type GameName string

const (
	EuroMillions GameName = "euromillions"
	Lotto        GameName = "lotto"
)

var EuroMillionsDays = map[Day]struct{}{
	Tuesday: {},
	Friday:  {},
}

var LottoDays = map[Day]struct{}{
	Wednesday: {},
	Saturday:  {},
}
