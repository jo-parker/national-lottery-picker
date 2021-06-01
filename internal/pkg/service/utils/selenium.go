package utils

import (
	"os"
	"github.com/tebeka/selenium"
)

func SaveScreenshot(wd selenium.WebDriver, path string) {
	data, err := wd.Screenshot()
	if err != nil {
		panic(err)
	}

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	f.Write(data)
}

func ClickElementByIDAndSendKeys(wd selenium.WebDriver, id string, text string) {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		panic(err)
	}
	elem.Click()
	elem.SendKeys(text)
}

func ClickElementByID(wd selenium.WebDriver, id string) {
	elem, err := wd.FindElement(selenium.ByID, id)
	if err != nil {
		SaveScreenshot(wd, "failure.png")
		panic(err)
	}
	elem.Click()
}

func ElementIsNotVisible(elt selenium.WebElement) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		visible, err := elt.IsDisplayed()
		return !visible, err
	}
}
