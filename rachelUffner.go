package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateRachelUffner(url string) {
	c := colly.NewCollector()
	var rachelUffnerExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-forthcoming_featured", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("span.date")
			location := "170 Suffolk Street New York, NY 10002"
			rachelUffnerExhibitions = append(rachelUffnerExhibitions, Exhibition{
				Gallery:   "Rachel Uffner Gallery",
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
	SaveToExcel(rachelUffnerExhibitions)
}
