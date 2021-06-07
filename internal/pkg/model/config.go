package model

import ()

type Config struct {
	NationalLottery struct {
		Game					GameName	`yaml:"game"`
		Draws					[]Draw		`yaml:"draws"`
		Username			string		`yaml:"username"`
		Password			string		`yaml:"password"`
		CostLimit			float32		`yaml:"costLimit"`
	}	`yaml:"national-lottery"`

	App struct {
		Debug						bool		`yaml:"debug"`
		Logfile					string	`yaml:"logfile"`
		ScreenshotDir		string	`yaml:"screenshotDir"`
	}	`yaml:"app"`
}
