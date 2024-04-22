package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateVsf(url string) {
	c := colly.NewCollector()
	var vsfExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-forthcoming", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("span.date")
			location := colElement.ChildText("")
			vsfExhibitions = append(vsfExhibitions, Exhibition{
				Gallery:   "vsf",
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
	SaveToExcel(vsfExhibitions)
}
