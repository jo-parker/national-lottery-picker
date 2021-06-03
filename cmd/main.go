package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
)

func main() {
	configPtr := flag.String("c", "/etc/national-lottery-picker/config.yml", "Configuration file path")
	flag.Parse()

	configFile, err := ioutil.ReadFile(*configPtr)
	if err != nil {
			panic(err)
	}

	var config model.Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
    panic(err)
	}

	service.Config = config
	service.EnterDraw()
}
