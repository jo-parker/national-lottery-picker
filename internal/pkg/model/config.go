package model

import ()

type Config struct {
	NationalLottery struct {
		Game					GameName	`yaml:"game"`
		Draws					[]Draw	`yaml:"draws"`
		Username			string	`yaml:"username"`
		Password			string	`yaml:"password"`
		CostLimit			float32	`yaml:"costLimit"`
		Cron					string	`yaml:"cron"`
	}	`yaml:"national-lottery"`

	App struct {
		Test						bool	`yaml:"test"`
		Logfile					string	`yaml:"logfile"`
		ScreenshotDir		string	`yaml:"screenshotDir"`
		BinDir					string	`yaml:"binDir"`
	}	`yaml:"app"`
}
