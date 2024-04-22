package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGreeneNaftali(url string) {
	c := colly.NewCollector()
	var greeneNaftaliExhibitions []Exhibition
	c.OnHTML("#upcoming", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.b.caps")
			title := colElement.ChildText("span.i")
			date := colElement.ChildText("time")
			location := colElement.ChildText("508 West 26th Street, Ground Floor & 8th Floor, New York, NY 10001")
			greeneNaftaliExhibitions = append(greeneNaftaliExhibitions, Exhibition{
				Gallery:   "Greene Naftali",
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
	SaveToExcel(greeneNaftaliExhibitions)
}
