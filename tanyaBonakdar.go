package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateTanyaBonakdar(url string) {
	c := colly.NewCollector()
	var tanyaBonakdarExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistAndTitle := colElement.ChildText("h2")
			artistAndTitleSplit := strings.Split(artistAndTitle, ": ")
			if len(artistAndTitleSplit) == 2 {
				artistText := artistAndTitleSplit[0]
				title := artistAndTitleSplit[1]
				date := colElement.ChildText("span.date")
				location := colElement.ChildText("span.subtitle")
				tanyaBonakdarExhibitions = append(tanyaBonakdarExhibitions, Exhibition{
					Gallery:   "Tanya Bonakdar",
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
	SaveToExcel(tanyaBonakdarExhibitions)
}
