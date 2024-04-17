package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateCanada(url string) {
	c := colly.NewCollector()
	var canadaExhibitions []Exhibition
	c.OnHTML("section", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			title := colElement.ChildText("h1")
			date := colElement.ChildText("time")
			artist := colElement.ChildText("p:nth-of-type(1)")
			location := colElement.ChildText("p:nth-of-type(3)")
			canadaExhibitions = append(canadaExhibitions, Exhibition{
				Gallery:   "Canada",
				Location:  location,
				Artist:    artist,
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
	SaveToExcel(canadaExhibitions)
}
