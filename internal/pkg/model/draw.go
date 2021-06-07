package model

import ()

type Draw struct {
	Name                GameName `yaml:"name"`
	NumTickets        	int			 `yaml:"numberOfTickets"`
	Day                 Day      `yaml:"day"`
}
