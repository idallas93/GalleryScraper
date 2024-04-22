package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGiselaCapitain(url string) {
	c := colly.NewCollector()
	var giselaCapitainExhibitions []Exhibition
	c.OnHTML("div.container", func(e *colly.HTMLElement) {
		e.ForEach("div.index-grid__text-wrapper", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h3")
			title := colElement.ChildText("p.index-grid__text-subtitle")
			date := colElement.ChildText("p.index-grid__text-date")
			location := colElement.ChildText("p.index-grid__text-location")
			giselaCapitainExhibitions = append(giselaCapitainExhibitions, Exhibition{
				Gallery:   "Gisela Capitain",
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
	SaveToExcel(giselaCapitainExhibitions)
}
