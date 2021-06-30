package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/jpparker/national-lottery-picker/internal/pkg/model"
)

var gameName model.GameName

const (
	Hot  NumberType = "hot"
	Cold NumberType = "cold"
)

type HotColdNumbers struct {
	Main    []int
	Special []int
}

type NumberType string

func HotColdScraper(gn model.GameName) (map[NumberType]HotColdNumbers, error) {
	gameName = gn
	url := fmt.Sprintf("https://www.national-lottery.com/%s/statistics", gameName)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Status code error: %d %s", res.StatusCode, res.Status))
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	numbers := map[NumberType]HotColdNumbers{}
	numbers[Hot] = findNumbers(doc, Hot)
	numbers[Cold] = findNumbers(doc, Cold)

	return numbers, nil
}

func findNumbers(doc *goquery.Document, numType NumberType) HotColdNumbers {
	numbers := HotColdNumbers{
		Main:    []int{},
		Special: []int{},
	}
	selector := fmt.Sprintf("h2:contains('%s Numbers') ~ table", numType)

	doc.Find(selector).Each(
		func(i int, s *goquery.Selection) {
			mainSelector := fmt.Sprintf(".%s.ball:not(.lucky-star)", gameName)
			s.Find(mainSelector).Each(
				func(i int, s *goquery.Selection) {
					val, _ := strconv.Atoi(s.Text())
					numbers.Main = append(numbers.Main, val)
				},
			)

			s.Find(".lucky-star").Each(
				func(i int, s *goquery.Selection) {
					val, _ := strconv.Atoi(s.Text())
					numbers.Special = append(numbers.Special, val)
				},
			)
		},
	)

	return numbers
}
