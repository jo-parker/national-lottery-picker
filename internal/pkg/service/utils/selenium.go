package utils

import (
	"os"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var Config model.Config

func SaveScreenshot(wd selenium.WebDriver, path string) {
	data, err := wd.Screenshot()
	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create(Config.App.ScreenshotDir + "/" + path)
	if err != nil {
		fmt.Println(err)
	}

	f.Write(data)
}

func ClickElementByIDAndSendKeys(wd selenium.WebDriver, id string, text string) error {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		return err
	}
	elem.Click()
	elem.SendKeys(text)

	return nil
}

func ClickElementByID(wd selenium.WebDriver, id string) error {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		return err
	}
	elem.Click()

	return nil
}
