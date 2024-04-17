package main

import (
	"github.com/gocolly/colly"
	"log"
)

func updateBartolami(url string) {
	c := colly.NewCollector()
	var bartolamiExhibitions []Exhibition
	c.OnHTML("#current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h4.order--2")
			title := colElement.ChildText("em")
			date := colElement.ChildText("h4.order--4")
			location := colElement.ChildText("h5")
			bartolamiExhibitions = append(bartolamiExhibitions, Exhibition{
				Gallery:   "Bartolami",
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
	SaveToExcel(bartolamiExhibitions)
}
