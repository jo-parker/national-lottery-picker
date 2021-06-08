package config

import (
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

type Config struct {
	NationalLottery struct {
		Game					model.GameName	`yaml:"game"`
		Draws					[]model.Draw		`yaml:"draws"`
		Username			string		`yaml:"username"`
		Password			string		`yaml:"password"`
		CostLimit			float32		`yaml:"costLimit"`
		Cron					string		`yaml:"cron"`
	}	`yaml:"national-lottery"`

	App struct {
		Debug						bool		`yaml:"debug"`
		Logfile					string	`yaml:"logfile"`
		ScreenshotDir		string	`yaml:"screenshotDir"`
	}	`yaml:"app"`
}
