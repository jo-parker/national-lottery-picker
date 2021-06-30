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
	"github.com/robfig/cron"
	"github.com/tebeka/selenium"
	"gopkg.in/yaml.v2"
)

var Config model.Config

func main() {
	configPtr := flag.String("c", "/etc/national-lottery-picker/config.yml", "Configuration file path")
	usernamePtr := flag.String("u", "username", "Username for the National Lottery account")
	passwordPtr := flag.String("p", "password", "Password for the National Lottery account")
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
	service.Username = *usernamePtr
	service.Password = *passwordPtr

	logfile, err := os.OpenFile(Config.App.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logfile)

	if err := initSelenium(); err != nil {
		log.Fatalln(err)
	}

	c := cron.New()
	c.AddFunc(Config.NationalLottery.Cron, func() {
		enterDraws()
	})
	c.Run()
}

func initSelenium() error {
	seleniumPath := fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", Config.App.BinDir)
	chromeDriverPath := fmt.Sprintf("%s/chromedriver-linux64", Config.App.BinDir)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	_, err := selenium.NewSeleniumService(seleniumPath, service.Port, opts...)
	if err != nil {
		return err
	}

	return nil
}

func enterDraws() {
	log.Println("[INFO] Entering draws...")

	for _, d := range Config.NationalLottery.Draws {
		log.Println(fmt.Sprintf("[INFO] Entering %s %s draw...", d.Name, d.Day))

		if d.NumTickets > 4 {
			log.Printf("[ERROR] Maximum number of 4 tickets exceeded in one order: %d", d.NumTickets)
			continue
		}

		if err := service.EnterDraw(d); err != nil {
			log.Printf("[ERROR] Entering draw failed: %s", err)
			continue
		}
	}

	log.Println("[INFO] Run complete.")
}
