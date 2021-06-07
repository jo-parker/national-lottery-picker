package main

import (
	"flag"
	"log"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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


	service.EnterDraw()
}
