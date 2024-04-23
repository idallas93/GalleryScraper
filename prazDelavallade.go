package main

import (
	"log"

	"github.com/gocolly/colly"
)

func updatePrazDelavallade(url string) {
	c := colly.NewCollector()
	var prazDelavalladeExhibitions []Exhibition
	c.OnHTML("div.main-content", func(e *colly.HTMLElement) {
		e.ForEach("div.summary-content.sqs-gallery-meta-container", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("p:nth-of-type(2)")
			date := colElement.ChildText("time:nth-of-type(1)")
			location := colElement.ChildText("p:nth-of-type(1)")
			prazDelavalladeExhibitions = append(prazDelavalladeExhibitions, Exhibition{
				Gallery:   "Praz Delavallade",
				Location:  location,
				Artist:    artistText,
				Title:     artistText,
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
	SaveToExcel(prazDelavalladeExhibitions)
}
