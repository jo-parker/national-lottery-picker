package model

import ()

type Ticket struct {
	MainNumbers     map[int]struct{}
	SpecialNumbers  map[int]struct{}
	Entered         bool
}

func (t *Ticket) Init() {
	t.MainNumbers = make(map[int]struct{})
	t.SpecialNumbers = make(map[int]struct{})
	t.Entered = false
}
