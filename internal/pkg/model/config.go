package model

import ()

type Config struct {
	NationalLottery struct {
		Draw            DrawName `yaml:"draw"`
		NumberOfTickets int      `yaml:"numberOfTickets"`
		Days            []Day 	 `yaml:"days"`
		Username        string   `yaml:"username"`
		Password        string   `yaml:"password"`
		CostLimit 			float32  `yaml:"costLimit"`
	} `yaml:"national-lottery"`

	App struct {
		Debug           bool     `yaml:"debug"`
		LogDir 					string 	 `yaml:"logdir"`
		ScreenshotDir   string   `yaml:"screenshotDir"
	} `yaml:"app"`
}
