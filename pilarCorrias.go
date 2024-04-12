package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updatePilarCorrias(url string) {
	c := colly.NewCollector()
	var pilarCorriasExhibitions []Exhibition
	c.OnHTML("#container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.heading_title")
			title := colElement.ChildText("div.subtitle.ani-in")
			date := colElement.ChildText("div.date")
			location := strings.ReplaceAll(date, colElement.ChildText("div.bottom.ani-in"), "")
			if artistText != "" {
				pilarCorriasExhibitions = append(pilarCorriasExhibitions, Exhibition{
					Gallery:   "Pilar Corrias",
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
	SaveToExcel(pilarCorriasExhibitions)
}
