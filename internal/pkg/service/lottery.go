package service

import (
	"os"
	"fmt"
	"log"
	"math"
	"strconv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	slog "github.com/tebeka/selenium/log"
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

func EnterDraw(draw *model.Draw) {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath), // Specify the path to ChromeWebDriver in order to use Chrome.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance.
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	loggingCaps := slog.Capabilities {
		slog.Server: slog.Info,
		slog.Browser: slog.Info,
		slog.Client: slog.Info,
		slog.Driver: slog.Info,
		slog.Performance: slog.Off,
		slog.Profiler: slog.Off,
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
		log.Fatalln(err)
	}
	defer wd.Quit()

	if err := wd.Get(baseUrl); err != nil {
		log.Fatalln(err)
	}

	consentCookie := &selenium.Cookie {
		Name: "CONSENTMGR",
		Value: "c1:0%7Cc3:0%7Cc9:0%7Cc11:0%7Cts:1622708767594%7Cconsent:false",
		Domain: ".national-lottery.co.uk",
		Path: "/",
		Expiry: math.MaxUint32,
	}
	if err := wd.AddCookie(consentCookie); err != nil {
		log.Fatalln(err)
	}

	if err := wd.Get(fmt.Sprintf("%s/sign-in", baseUrl)); err != nil {
		log.Fatalln(err)
	}
	utils.ClickElementByIDAndSendKeys(wd, "form_username", Config.NationalLottery.Username)
	utils.ClickElementByIDAndSendKeys(wd, "form_password", Config.NationalLottery.Password)
	utils.ClickElementByID(wd, "login_submit_bttn")

	playGame(wd, draw)

	utils.SaveScreenshot(wd, fmt.Sprintf("%s_%s_success.png", draw.Name, draw.Day))
}

func playGame(wd selenium.WebDriver, d *model.Draw) {
	var url string
	var gameDays map[model.Day]struct{}

	switch d.Name {
	case model.EuroMillions:
		url = fmt.Sprintf("%s/games/euromillions?icid=-:mm:-:mdg:em:dbg:pl:co", baseUrl)
		gameDays = model.EuroMillionsDays
	case model.Lotto:
		url = fmt.Sprintf("%s/games/lotto?icid=-:mm:-:mdg:lo:dbg:pl:co", baseUrl)
		gameDays = model.LottoDays
	}

	if err := wd.Get(url); err != nil {
		log.Fatalln(err)
	}

	populateTickets(wd, d)

	if _, ok := gameDays[d.Day]; ok {
		id := fmt.Sprintf("%s_dd_label", d.Day)
		utils.ClickElementByID(wd, id)
	} else {
		log.Fatalln(d.Name + " is not played on this day, exiting.")
	}

	utils.ClickElementByID(wd, fmt.Sprintf("%s_playslip_confirm", d.Name))

	placeOrder(wd)
}

func populateTickets(wd selenium.WebDriver, d *model.Draw) {
	for i := 0; i < d.NumTickets; i++ {
		utils.ClickElementByID(wd, fmt.Sprintf("number_picker_initialiser_%d", i))

		t := GenerateTicket(d)

		for key := range t.MainNumbers {
			utils.ClickElementByID(wd, fmt.Sprintf("pool_0_label_ball_%d", key))
		}
		for key := range t.SpecialNumbers {
			utils.ClickElementByID(wd, fmt.Sprintf("pool_1_label_ball_%d", key))
		}
		utils.ClickElementByID(wd, "number_selection_confirm_button")
	}

	if _, err := wd.ExecuteScript("document.querySelector('label#weeks1',':before').click();", nil); err != nil {
	  utils.SaveScreenshot(wd, "failure.png")
		log.Fatalln(err)
	}
}

func placeOrder(wd selenium.WebDriver) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, "span#price")
	if err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		log.Fatalln(err)
	}
	price, err := elem.GetAttribute("data-price")
	if err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		log.Fatalln(err)
	}

	priceFloat, _ := strconv.ParseFloat(price, 32)

	costLimitExceeded := float32(priceFloat) > Config.NationalLottery.CostLimit
	if costLimitExceeded {
		utils.SaveScreenshot(wd, "failure.png")
		log.Fatalln("Configured cost limit exceeded when reviewing order, saving screenshot")
	}

	// Place order
	if !Config.App.Debug {
		utils.ClickElementByID(wd, "confirm")
	}
}
