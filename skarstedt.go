package main

import (
	"github.com/gocolly/colly"
	"log"
)

func updateSkarstedt(url string) {
	c := colly.NewCollector()
	var skarstedtExhibitions []Exhibition
	c.OnHTML("#exhibitions-container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("3")
			location := colElement.ChildText("div.list-detail-title")
			skarstedtExhibitions = append(skarstedtExhibitions, Exhibition{
				Gallery:   "Skarstedt",
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
	SaveToExcel(skarstedtExhibitions)
}
