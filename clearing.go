package main

import (
	"github.com/gocolly/colly"
	"log"
)

func updateClearing(url string) {
	c := colly.NewCollector()
	var clearingExhibitions []Exhibition
	c.OnHTML("#contenair", func(e *colly.HTMLElement) {
		e.ForEach("section:nth-of-type(1)", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("p:nth-of-type(1)")
			title := ""
			date := colElement.ChildText("p:nth-of-type(2)")
			location := colElement.ChildText("p:nth-of-type(3)")
			clearingExhibitions = append(clearingExhibitions, Exhibition{
				Gallery:   "Clearing",
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
	SaveToExcel(clearingExhibitions)
}
