package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updatePpow(url string) {
	c := colly.NewCollector()
	var ppowExhibitions []Exhibition
	c.OnHTML("#medium", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2.subtitle2")
			title = strings.ReplaceAll(title, location, "")
			ppowExhibitions = append(ppowExhibitions, Exhibition{
				Gallery:   "P.P.O.W.",
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
	SaveToExcel(ppowExhibitions)
}
