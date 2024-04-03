package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateLehmannMaupin(url string) {
	c := colly.NewCollector()
	var lehmannMaupinExhibitions []Exhibition
	c.OnHTML("#exhibitions-container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2.subtitle2")
			lehmannMaupinExhibitions = append(lehmannMaupinExhibitions, Exhibition{
				Gallery:   "Lehmann Maupin",
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
	SaveToExcel(lehmannMaupinExhibitions)
}
