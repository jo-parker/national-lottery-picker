package model

import ()

type Config struct {
	NationalLottery struct {
		Draw            DrawName `yaml:"draw"`
		NumberOfTickets int      `yaml:"numberOfTickets"`
		Days            []string `yaml:"days"`
		Username        string   `yaml:"username"`
		Password        string   `yaml:"password"`
	} `yaml:"national-lottery"`

	App struct {
		Debug           bool     `yaml:"debug"`
		LogDir 					string 	 `yaml:"logdir"`
	} `yaml:"app"`
}
