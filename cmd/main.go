package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/robfig/cron"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
)

func main() {
	configPtr := flag.String("c", "/etc/national-lottery-picker/config.yml", "Configuration file path")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configPtr)
	if err != nil {
		log.Fatalln(err)
	}

	var config model.Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	service.Config = config
	utils.Config = config

	logfile, err := os.OpenFile(config.App.Logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(logfile)

	c := cron.New()
	c.AddFunc(config.NationalLottery.Cron, func() {
		log.Println("Entering draws...")
		for _, d := range config.NationalLottery.Draws {
			log.Println(fmt.Sprintf("Entering %s %s draw...", d.Name, d.Day))

			if d.NumTickets > 4 {
				log.Printf("Maximum number of 4 tickets exceeded in one order: ", d.NumTickets)
				continue
			}

			if err := service.EnterDraw(&d); err != nil {
				log.Printf("Entering draw failed: ", err)
				continue
			}
		}
		log.Println("Run complete.")
	})
	c.Run()
}
