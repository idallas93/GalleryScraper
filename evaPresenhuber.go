package main

import (
	"github.com/gocolly/colly"
	"log"
)

func updateEvaPresenhuber(url string) {
	c := colly.NewCollector()
	var evaPresenhuberExhibitions []Exhibition
	c.OnHTML("#medium", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2:nth-of-type(1)")
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2:nth-of-type(2)")
			evaPresenhuberExhibitions = append(evaPresenhuberExhibitions, Exhibition{
				Gallery:   "Eva Presenhuber",
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
	SaveToExcel(evaPresenhuberExhibitions)
}
