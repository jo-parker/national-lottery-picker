package service

import (
	"os"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/jpparker/euromillions-picker/internal/pkg/model"
	"github.com/serge1peshcoff/selenium-go-conditions"
)

const (
	vendorPath      = "/app/euromillions-picker/vendor"
	port            = 8080
	baseUrl         = "https://www.national-lottery.co.uk/"
)

var seleniumPath = fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", vendorPath)
var geckoDriverPath = fmt.Sprintf("%s/geckodriver-v0.29.1-linux64", vendorPath)

func PlayEuromillions() {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	//selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get(fmt.Sprintf("%s/sign-in", baseUrl)); err != nil {
		panic(err)
	}

	// Accept Cookie consent
	if err = wd.Wait(conditions.ElementIsLocated(selenium.ByCSSSelector, ".cuk_btn_primary:nth-child(2)")); err != nil {
		panic(err)
	}
	elem, err := wd.FindElement(selenium.ByCSSSelector, ".cuk_btn_primary:nth-child(2)")
	if err != nil {
		panic(err)
	}
	elem.Click(); elem.Click(); elem.Click()

	elem, _ = wd.FindElement(selenium.ByCSSSelector, ".cu_k_modal_main_box")
	if err = wd.Wait(ElementIsNotVisible(elem)); err != nil {
		panic(err)
	}

	SignIn(wd)

	entry := GenerateTicket()
	EnterDraw(wd, entry)

	SaveScreenshot(wd, "screenshot.png")
}

func EnterDraw(wd selenium.WebDriver, ticket model.Ticket) {
	// Go to Euromillions game
	if err := wd.Get(fmt.Sprintf("%s/games/euromillions?icid=-:mm:-:mdg:em:dbg:pl:co", baseUrl)); err != nil {
		panic(err)
	}
}

func SignIn(wd selenium.WebDriver) {
	elem, err := wd.FindElement(selenium.ByID, "form_username")
	if err != nil {
		panic(err)
	}
	elem.Click()
	elem.SendKeys("jpparker3986@hotmail.co.uk")

	elem, err = wd.FindElement(selenium.ByID, "form_password")
	if err != nil {
		panic(err)
	}
	elem.Click()
	elem.SendKeys("Agj4xaaX")

	elem, err = wd.FindElement(selenium.ByID, "login_submit_bttn")
	if err != nil {
		panic(err)
	}
	elem.Click()
}

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

func ElementIsNotVisible(elt selenium.WebElement) selenium.Condition {
	return func(wd selenium.WebDriver) (bool, error) {
		visible, err := elt.IsDisplayed()
		return !visible, err
	}
}
