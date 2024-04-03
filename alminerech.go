package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateAlmineRech(url string) {
	c := colly.NewCollector()
	var almineRechExhibitions []Exhibition
	c.OnHTML("#main", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.title")
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("div.info")
			location := colElement.ChildText("div.tag")
			if artistText != "" {
				almineRechExhibitions = append(almineRechExhibitions, Exhibition{
					Gallery:   "Almine Rech",
					Location:  location,
					Artist:    artistText,
					Title:     title,
					StartDate: date,
					EndDate:   date,
					Notes:     "",
				})
			} 
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(almineRechExhibitions)
}
