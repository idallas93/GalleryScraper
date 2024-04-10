package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updatePerrotin(url string) {
	c := colly.NewCollector()
	var perrotinExhibitions []Exhibition
	c.OnHTML("#listing", func(e *colly.HTMLElement) {
		e.ForEach("div.infos", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("p.fblack.fsmall")
			title := colElement.ChildText("p.fblack.flarge")
			date := colElement.ChildText("p.fblack.fsmall.date")
			address := colElement.ChildText("p.fbold.flsmall.address.hidden")
			if artistText != "" {
				perrotinExhibitions = append(perrotinExhibitions, Exhibition{
					Gallery:   "Perrotin Gallery",
					Location:  address,
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
	SaveToExcel(perrotinExhibitions)
}
