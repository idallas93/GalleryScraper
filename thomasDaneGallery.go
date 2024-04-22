package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateThomasDaneGallery(url string) {
	c := colly.NewCollector()
	var thomasDaneGalleryexhibitions []Exhibition
	c.OnHTML("ul.current-forthcoming", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h3")
			title := colElement.ChildText("span.title")
			date := colElement.ChildText("span.date")
			location := colElement.ChildText("span.location")
			thomasDaneGalleryexhibitions = append(thomasDaneGalleryexhibitions, Exhibition{
				Gallery:   "Thomas Dane Gallery",
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
	SaveToExcel(thomasDaneGalleryexhibitions)
}
