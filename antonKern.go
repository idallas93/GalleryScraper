package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateAntonKern(url string) {
	c := colly.NewCollector()
	var antonKernExhibitions []Exhibition
	c.OnHTML("#main_content", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.ChildText("h2")
			splitAllText := strings.Split(allText, ":")
			var artistText string
			var title string
			if len(splitAllText) == 2 {
				artistText = splitAllText[0]
				title = splitAllText[1]
			}
			date := colElement.ChildText("span.date")
			readMore := colElement.ChildText("span.read_more")
			location := colElement.ChildText("span.subtitle")
			if readMore != "" {
				antonKernExhibitions = append(antonKernExhibitions, Exhibition{
					Gallery:   "Anton Kern",
					Location:  location,
					Artist:    artistText,
					Title:     title,
					StartDate: date,
					EndDate:   date,
					Notes:     "",
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(antonKernExhibitions)
}
