package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jpparker/national-lottery-picker/internal/pkg/handlers"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service"
	"github.com/jpparker/national-lottery-picker/internal/pkg/service/utils"
	"github.com/tebeka/selenium"
	"gopkg.in/yaml.v2"
)

var Config model.Config

func init() {
	configFile, err := ioutil.ReadFile("/etc/national-lottery-picker/config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(configFile, &Config)
	if err != nil {
		log.Fatalln(err)
	}

	service.Config = Config
	utils.Config = Config

	seleniumPath := fmt.Sprintf("%s/selenium-server-standalone-3.141.59.jar", Config.App.BinDir)
	chromeDriverPath := fmt.Sprintf("%s/chromedriver-linux64", Config.App.BinDir)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	_, err = selenium.NewSeleniumService(seleniumPath, service.Port, opts...)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return handlers.EnterDraws(req)
	default:
		return handlers.UnhandledMethod()
	}
}
