package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateLuhringAugustine(url string) {
	c := colly.NewCollector()
	var luhringAugustineExhibitions []Exhibition
	c.OnHTML("#exhibitions-container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			location := colElement.ChildText("h2.subtitle2")
			title := strings.ReplaceAll(colElement.ChildText("h2"), location, "")
			date := colElement.ChildText("span.custom-br")
			luhringAugustineExhibitions = append(luhringAugustineExhibitions, Exhibition{
				Gallery:   "Luhring Augustine",
				Location:  location,
				Artist:    artistText,
				Title:     title,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(luhringAugustineExhibitions)
}
