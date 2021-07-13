package model

type Config struct {
	App struct {
		Test          bool   `yaml:"test"`
		Logfile       string `yaml:"logfile"`
		ScreenshotDir string `yaml:"screenshotDir"`
		BinDir        string `yaml:"binDir"`
	} `yaml:"app"`
}
