package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

func updateLissonData(url string) {
	c := colly.NewCollector()

	var lissonData []Exhibition
	c.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach("div.news-grid-3up-item", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.Text
			space := regexp.MustCompile(`\s{2,}`)
			allText = space.ReplaceAllString(allText, ";")
			splitText := strings.Split(allText, ";")
			textWithoutEmptyVals := delete_empty(splitText)
			if len(textWithoutEmptyVals) == 3 {
				artistText := strings.Split(textWithoutEmptyVals[0], ":")
				if len(artistText) == 2 {
					lissonData = append(lissonData, Exhibition{
						Gallery:   "Lisson",
						Location:  textWithoutEmptyVals[2],
						Artist:    artistText[0],
						Title:     artistText[1],
						StartDate: textWithoutEmptyVals[1],
						EndDate:   textWithoutEmptyVals[1],
						Notes:     "",
					})
				} else {
					lissonData = append(lissonData, Exhibition{
						Gallery:   "Lisson",
						Location:  textWithoutEmptyVals[2],
						Artist:    textWithoutEmptyVals[0],
						Title:     "Untitled",
						StartDate: textWithoutEmptyVals[1],
						EndDate:   textWithoutEmptyVals[1],
						Notes:     "",
					})
				}
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(lissonData)
}

func delete_empty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}
