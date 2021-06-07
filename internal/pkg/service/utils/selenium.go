package utils

import (
	"os"
	"log"
	"github.com/tebeka/selenium"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var Config model.Config

func SaveScreenshot(wd selenium.WebDriver, path string) {
	data, err := wd.Screenshot()
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(Config.App.ScreenshotDir + "/" + path)
	if err != nil {
		log.Fatalln(err)
	}

	f.Write(data)
}

func ClickElementByIDAndSendKeys(wd selenium.WebDriver, id string, text string) {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		log.Fatalln(err)
	}
	elem.Click()
	elem.SendKeys(text)
}

func ClickElementByID(wd selenium.WebDriver, id string) {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		log.Fatalln(err)
	}
	elem.Click()
}
