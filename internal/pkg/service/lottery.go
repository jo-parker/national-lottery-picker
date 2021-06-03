package service

import (
	"os"
	"fmt"
	"math"
	"strconv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/log"
	"github.com/tebeka/selenium/chrome"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

const (
	vendorPath      = "/app/national-lottery-picker/vendor"
	port            = 8080
	baseUrl         = "https://national-lottery.co.uk/"
)

var Config model.Config
var seleniumPath = fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", vendorPath)
var chromeDriverPath = fmt.Sprintf("%s/chromedriver-linux64", vendorPath)

func EnterDraw() {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	loggingCaps := log.Capabilities {
		log.Server: log.Severe,
		log.Browser: log.Severe,
		log.Client: log.Severe,
		log.Driver: log.Severe,
		log.Performance: log.Off,
		log.Profiler: log.Off,
	}
	chromeCaps := chrome.Capabilities {
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--log-level=2",
		},
	}
	caps.AddLogging(loggingCaps)
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get(baseUrl); err != nil {
					panic(err)
	}

	consentCookie := &selenium.Cookie {
		Name: "CONSENTMGR",
		Value: "c1:0%7Cc3:0%7Cc9:0%7Cc11:0%7Cts:1622708767594%7Cconsent:false",
		Domain: ".national-lottery.co.uk",
		Path: "/",
		Expiry: math.MaxUint32,
	}
	if err := wd.AddCookie(consentCookie); err != nil {
					panic(err)
	}

	if err := wd.Get(fmt.Sprintf("%s/sign-in", baseUrl)); err != nil {
					panic(err)
	}
	utils.ClickElementByIDAndSendKeys(wd, "form_username", Config.NationalLottery.Username)
	utils.ClickElementByIDAndSendKeys(wd, "form_password", Config.NationalLottery.Password)
	utils.ClickElementByID(wd, "login_submit_bttn")

	t := GenerateTicket()
	switch t.Draw.Name {
	case model.Euromillions:
		playEuromillions(wd, t)
	case model.Lotto:
		playLotto(wd, t)
	}

	utils.SaveScreenshot(wd, "success.png")
}

func playEuromillions(wd selenium.WebDriver, t model.Ticket) {
	if err := wd.Get(fmt.Sprintf("%s/games/euromillions?icid=-:mm:-:mdg:em:dbg:pl:co", baseUrl)); err != nil {
		panic(err)
	}

	// Populate ticket
	utils.ClickElementByID(wd, "number_picker_initialiser_0")

	for key := range t.MainNumbers {
		utils.ClickElementByID(wd, fmt.Sprintf("pool_0_label_ball_%d", key))
	}

	for key := range t.SpecialNumbers {
		utils.ClickElementByID(wd, fmt.Sprintf("pool_1_label_ball_%d", key))
	}

	utils.ClickElementByID(wd, "number_selection_confirm_button")
	utils.ClickElementByID(wd, "fri_dd_label")
	if _, err := wd.ExecuteScript("document.querySelector('label#weeks1',':before').click();", nil); err != nil {
		panic(err)
	}
	utils.ClickElementByID(wd, "euromillions_playslip_confirm")

	placeOrder(wd)
}

func placeOrder(wd selenium.WebDriver) {
	// Check cost threshold
	elem, err := wd.FindElement(selenium.ByCSSSelector, "span#price")
	if err != nil {
		panic(err)
	}
	price, err := elem.GetAttribute("data-price")
	priceFloat, _ := strconv.ParseFloat(price, 32)
	if priceFloat > 5.0 || err != nil {
		panic(err)
	}

	// Place order
	if !Config.App.Debug {
		utils.ClickElementByID(wd, "confirm")
	}
}

func playLotto(wd selenium.WebDriver, t model.Ticket) {
}
