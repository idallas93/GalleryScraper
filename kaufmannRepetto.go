package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateKaufmannRepetto(url string) {
	c := colly.NewCollector()
	var kaufmannRepettoExhibitions []Exhibition
	c.OnHTML("div.main-container", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("strong.author")
			title := colElement.ChildText("em.opera")
			allText := colElement.ChildText("p")
			var location string
			var date string
			if strings.Contains(allText, "through") {
				splitThroughText := strings.Split(allText, "through")
				location = splitThroughText[0]
				date = splitThroughText[1]
			} else if strings.Contains(allText, "opening") {
				splitThroughText := strings.Split(allText, "opening")
				location = splitThroughText[0]
				date = splitThroughText[1]
			}
			if artistText != "" {
				kaufmannRepettoExhibitions = append(kaufmannRepettoExhibitions, Exhibition{
					Gallery:   "Kaufmann Repetto",
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
	SaveToExcel(kaufmannRepettoExhibitions)
}
