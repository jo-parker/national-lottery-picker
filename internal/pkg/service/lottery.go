package service

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	slog "github.com/tebeka/selenium/log"
)

//go:generate mockgen -source lottery.go -destination mocks/lottery.go
type Lottery interface {
	EnterDraws(draws []model.Draw, credentials model.Credentials) []error
}

type LotteryImpl struct{}

const (
	Port    = 8080
	baseUrl = "https://national-lottery.co.uk/"
)

var Config model.Config
var caps selenium.Capabilities

func init() {
	caps = selenium.Capabilities{
		"browserName": "chrome",
	}
	loggingCaps := slog.Capabilities{
		slog.Server:      slog.Info,
		slog.Browser:     slog.Info,
		slog.Client:      slog.Info,
		slog.Driver:      slog.Info,
		slog.Performance: slog.Off,
		slog.Profiler:    slog.Off,
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
			"--no-sandbox",
			"--log-level=2",
		},
	}
	caps.AddLogging(loggingCaps)
	caps.AddChrome(chromeCaps)
}

func (impl *LotteryImpl) EnterDraws(draws []model.Draw, credentials model.Credentials) []error {
	var errors []error

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", Port))
	if err != nil {
		log.Fatalln(err)
	}
	defer wd.Quit()

	if err := wd.Get(baseUrl); err != nil {
		log.Fatalln(err)
	}

	consentCookie := &selenium.Cookie{
		Name:   "CONSENTMGR",
		Value:  "c1:0%7Cc3:0%7Cc9:0%7Cc11:0%7Cts:1622708767594%7Cconsent:false",
		Domain: ".national-lottery.co.uk",
		Path:   "/",
		Expiry: math.MaxUint32,
	}
	if err := wd.AddCookie(consentCookie); err != nil {
		log.Fatalln(err)
	}

	// sign-in user
	if err := wd.Get(fmt.Sprintf("%s/sign-in", baseUrl)); err != nil {
		log.Fatalln(err)
	}
	if err := utils.ClickElementByIDAndSendKeys(wd, "form_username", credentials.Username); err != nil {
		return append(errors, err)
	}
	if err := utils.ClickElementByIDAndSendKeys(wd, "form_password", credentials.Password); err != nil {
		return append(errors, err)
	}
	if err := utils.ClickElementByID(wd, "login_submit_bttn"); err != nil {
		return append(errors, err)
	}
	if url, _ := wd.CurrentURL(); strings.Contains(url, "login_error=1") {
		err := fmt.Errorf("login credentials for %s are incorrect", credentials.Username)
		return append(errors, err)
	}

	for _, d := range draws {
		if d.NumTickets > 4 {
			errors = append(errors, fmt.Errorf("maximum number of 4 tickets exceeded in one order: %d", d.NumTickets))
			continue
		}

		if err := enterDraw(wd, d, credentials); err != nil {
			errors = append(errors, fmt.Errorf("entering draw failed: %s", err))
			continue
		}
	}

	return errors
}

func enterDraw(wd selenium.WebDriver, draw model.Draw, credentials model.Credentials) error {
	var url string
	var gameDays map[model.Day]struct{}

	switch draw.Name {
	case model.EuroMillions:
		url = fmt.Sprintf("%s/games/euromillions?icid=-:mm:-:mdg:em:dbg:pl:co", baseUrl)
		gameDays = model.EuroMillionsDays
	case model.Lotto:
		url = fmt.Sprintf("%s/games/lotto?icid=-:mm:-:mdg:lo:dbg:pl:co", baseUrl)
		gameDays = model.LottoDays
	default:
		return fmt.Errorf("unknown game '%s'", draw.Name)
	}

	if err := wd.Get(url); err != nil {
		return err
	}

	populateTickets(wd, draw)

	if _, exists := gameDays[draw.Day]; exists {
		id := fmt.Sprintf("%s_dd_label", draw.Day)
		utils.ClickElementByID(wd, id)
	} else {
		return fmt.Errorf("%s is not played on this day, exiting", draw.Name)
	}

	if err := utils.ClickElementByID(wd, fmt.Sprintf("%s_playslip_confirm", draw.Name)); err != nil {
		return err
	}

	// place order if debug disabled
	if !Config.App.Debug {
		if err := utils.ClickElementByID(wd, "confirm"); err != nil {
			return err
		}
	}

	utils.SaveScreenshot(wd, fmt.Sprintf("%s_%s_success.png", draw.Name, draw.Day))

	return nil
}

func populateTickets(wd selenium.WebDriver, d model.Draw) error {
	for i := 0; i < d.NumTickets; i++ {
		if err := utils.ClickElementByID(wd, fmt.Sprintf("number_picker_initialiser_%d", i)); err != nil {
			return err
		}

		t, err := GenerateTicket(d)
		if err != nil {
			return err
		}

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

		log.Println(fmt.Sprintf("[INFO] Ticket confirmed: %d, %d", t.MainNumbers, t.SpecialNumbers))
	}

	if _, err := wd.ExecuteScript("document.querySelector('label#weeks1',':before').click();", nil); err != nil {
		utils.SaveScreenshot(wd, "failure.png")
		return err
	}

	return nil
}
