package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGrimm(url string) {
	c := colly.NewCollector()
	var grimmExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-container", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.subtitle")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("span.date")
			location := colElement.ChildText("span.location")
			if location != "" {
				grimmExhibitions = append(grimmExhibitions, Exhibition{
					Gallery:   "Grimm",
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
	SaveToExcel(grimmExhibitions)
}
