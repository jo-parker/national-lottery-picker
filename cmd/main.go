package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/tebeka/selenium"
	"gopkg.in/yaml.v2"
)

var Config model.Config

func init() {
	seleniumPath := fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", Config.App.BinDir)
	chromeDriverPath := fmt.Sprintf("%s/chromedriver-linux64", Config.App.BinDir)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	_, err := selenium.NewSeleniumService(seleniumPath, service.Port, opts...)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	configPtr := flag.String("c", "/etc/national-lottery-picker/config.yml", "Configuration file path")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configPtr)
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	service.Config = Config
	utils.Config = Config

	logfile, err := os.OpenFile(Config.App.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logfile)
}
