package model

type Config struct {
	App struct {
		Debug         bool   `yaml:"debug"`
		Logfile       string `yaml:"logfile"`
		ScreenshotDir string `yaml:"screenshotDir"`
		BinDir        string `yaml:"binDir"`
	} `yaml:"app"`
}
