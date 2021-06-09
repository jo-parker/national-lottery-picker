package service

import (
	"fmt"
	"log"
	"math"
	"errors"
	"strconv"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	slog "github.com/tebeka/selenium/log"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

const (
	Port			= 8080
	baseUrl		= "https://national-lottery.co.uk/"
)

var Config model.Config
var Username, Password string

func EnterDraw(draw *model.Draw) error {

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

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", Port))
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

	if err := utils.ClickElementByIDAndSendKeys(wd, "form_username", Username); err != nil {
		return err
	}
	if err := utils.ClickElementByIDAndSendKeys(wd, "form_password", Password); err != nil {
		return err
	}
	if err := utils.ClickElementByID(wd, "login_submit_bttn"); err != nil {
		return err
	}

	if err := playGame(wd, draw); err != nil {
		return err
	}

	utils.SaveScreenshot(wd, fmt.Sprintf("%s_%s_success.png", draw.Name, draw.Day))

	return nil
}

func playGame(wd selenium.WebDriver, d *model.Draw) error {
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
		return err
	}

	populateTickets(wd, d)

	if _, exists := gameDays[d.Day]; exists {
		id := fmt.Sprintf("%s_dd_label", d.Day)
		utils.ClickElementByID(wd, id)
	} else {
		return errors.New(fmt.Sprintf("%s is not played on this day, exiting.", d.Name))
	}

	if err := utils.ClickElementByID(wd, fmt.Sprintf("%s_playslip_confirm", d.Name)); err != nil {
		return err
	}

	if err := placeOrder(wd); err != nil {
		return err
	}

	return nil
}

func populateTickets(wd selenium.WebDriver, d *model.Draw) error {
	for i := 0; i < d.NumTickets; i++ {
		if err := utils.ClickElementByID(wd, fmt.Sprintf("number_picker_initialiser_%d", i)); err != nil {
			return err
		}

		t := GenerateTicket(d)

		for key := range t.MainNumbers {
			if err := utils.ClickElementByID(wd, fmt.Sprintf("pool_0_label_ball_%d", key)); err != nil {
				return err
			}
		}
		for key := range t.SpecialNumbers {
			if err := utils.ClickElementByID(wd, fmt.Sprintf("pool_1_label_ball_%d", key)); err != nil {
				return err
			}
		}

		if err := utils.ClickElementByID(wd, "number_selection_confirm_button"); err != nil {
			return err
		}

		log.Println(fmt.Sprintf("Ticket confirmed: %d, %d", t.MainNumbers, t.SpecialNumbers))
	}

	if _, err := wd.ExecuteScript("document.querySelector('label#weeks1',':before').click();", nil); err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		return err
	}

	return nil
}

func placeOrder(wd selenium.WebDriver) error {
	elem, err := wd.FindElement(selenium.ByCSSSelector, "span#price")
	if err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		return err
	}
	price, err := elem.GetAttribute("data-price")
	if err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		return err
	}

	priceFloat, _ := strconv.ParseFloat(price, 32)

	costLimitExceeded := float32(priceFloat) > Config.NationalLottery.CostLimit
	if costLimitExceeded {
		utils.SaveScreenshot(wd, "failure.png")
		return errors.New("Configured cost limit exceeded when reviewing order, saving screenshot")
	}

	// Place order
	if !Config.App.Debug {
		if err := utils.ClickElementByID(wd, "confirm"); err != nil {
			return err
		}
	}

	return nil
}
