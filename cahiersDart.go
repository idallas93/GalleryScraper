package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateCahiersDart(url string) {
	c := colly.NewCollector()
	var cahiersDartExhibitions []Exhibition
	c.OnHTML("div.content", func(e *colly.HTMLElement) {
			artist := e.ChildText("h1:nth-of-type(1)")
			title := e.ChildText("h1:nth-of-type(2)")
			date := e.ChildText("h3")
			cahiersDartExhibitions = append(cahiersDartExhibitions, Exhibition{
				Gallery:   "Cahiers d'Art",
				Location:  "Éditions Cahiers d’Art 14–15 rue du Dragon 75006 Paris",
				Artist:    artist,
				Title:     title,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(cahiersDartExhibitions)
}
