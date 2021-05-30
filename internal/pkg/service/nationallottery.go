package service

import (
	"os"
	"fmt"
	"github.com/tebeka/selenium"
)

const (
	// These paths will be different on your system.
	vendorPath 			= "/app/euromillions-picker/vendor"
	port            = 8080
)

var seleniumPath = fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", vendorPath)
var geckoDriverPath = fmt.Sprintf("%s/geckodriver-v0.29.1-linux64", vendorPath)

func Test() {
	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}
	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()
}