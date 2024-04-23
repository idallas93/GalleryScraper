package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateSeanKelly(url string) {
	c := colly.NewCollector()
	var seanKellyExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2:nth-of-type(1)")
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2:nth-of-type(2)")
			seanKellyExhibitions = append(seanKellyExhibitions, Exhibition{
				Gallery:   "Sean Kelly",
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
	SaveToExcel(seanKellyExhibitions)
}
