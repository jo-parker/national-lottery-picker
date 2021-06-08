package model

import ()

type Draw struct {
	Name					GameName	`yaml:"name"`
	NumTickets		int				`yaml:"numberOfTickets"`
	Day						Day				`yaml:"day"`
}

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
