package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateKurimanzutto(url string) {
	c := colly.NewCollector()
	var kurimanzuttoExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := ""
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2")
			kurimanzuttoExhibitions = append(kurimanzuttoExhibitions, Exhibition{
				Gallery:   "Kurimanzutto",
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
	SaveToExcel(kurimanzuttoExhibitions)
}
