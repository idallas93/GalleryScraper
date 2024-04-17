package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateCaseyKaplan(url string) {
	c := colly.NewCollector()
	var caseyKaplanExhibitions []Exhibition
	c.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach("div.top-item", func(i int, colElement *colly.HTMLElement) {
			title := colElement.ChildText("h2:nth-of-type(1)")
			artist := colElement.ChildText("h2:nth-of-type(2)")
			date := colElement.ChildText("h2:nth-of-type(3)")
			location := "121 WEST 27TH STREET NEW YORK, NY 10001"
			caseyKaplanExhibitions = append(caseyKaplanExhibitions, Exhibition{
				Gallery:   "Casey Kaplan",
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
	SaveToExcel(caseyKaplanExhibitions)
}
