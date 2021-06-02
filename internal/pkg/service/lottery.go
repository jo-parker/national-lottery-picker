package service

import (
	"os"
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/serge1peshcoff/selenium-go-conditions"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

const (
	vendorPath      = "/app/national-lottery-picker/vendor"
	port            = 8080
	baseUrl         = "https://www.national-lottery.co.uk/"
)

var seleniumPath = fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", vendorPath)
var geckoDriverPath = fmt.Sprintf("%s/geckodriver-v0.29.1-linux64", vendorPath)

func EnterDraw(draw model.DrawName) {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

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

	// Navigate to the sign-in page.
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
	// Accept Cookies pop-up requires multiple clicks
	elem.Click(); elem.Click(); elem.Click()

	elem, _ = wd.FindElement(selenium.ByCSSSelector, ".cu_k_modal_main_box")
	if err = wd.Wait(utils.ElementIsNotVisible(elem)); err != nil {
		panic(err)
	}

	SignIn(wd)

	t := GenerateTicket(draw)
	switch t.Draw.Name {
	case model.Euromillions:
		PlayEuromillions(wd, t)
	}

	utils.SaveScreenshot(wd, "screenshot.png")
}

func PlayEuromillions(wd selenium.WebDriver, t model.Ticket) {
	if err := wd.Get(fmt.Sprintf("%s/games/euromillions?icid=-:mm:-:mdg:em:dbg:pl:co", baseUrl)); err != nil {
		panic(err)
	}

	utils.ClickElementByID(wd, "number_picker_initialiser_0")

	for key := range t.MainNumbers {
		utils.ClickElementByID(wd, fmt.Sprintf("pool_0_label_ball_%d", key))
	}

	for key := range t.SpecialNumbers {
		utils.ClickElementByID(wd, fmt.Sprintf("pool_1_label_ball_%d", key))
	}

	utils.ClickElementByID(wd, "number_selection_confirm_button")
	utils.ClickElementByID(wd, "tue_dd_label")
	utils.ClickElementByID(wd, "weeks1")
	utils.ClickElementByID(wd, "euromillions_playslip_confirm")
}

func SignIn(wd selenium.WebDriver) {
	utils.ClickElementByIDAndSendKeys(wd, "form_username", "jpparker3986@hotmail.co.uk")
	utils.ClickElementByIDAndSendKeys(wd, "form_password", "Agj4xaaX")
	utils.ClickElementByID(wd, "login_submit_bttn")
}
