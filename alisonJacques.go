package main

import (
	"github.com/gocolly/colly"
	"log"
)

func updateAlisonJacques(url string) {
	c := colly.NewCollector()
	var paceExhibitions []Exhibition
	c.OnHTML("#main", func(e *colly.HTMLElement) {
		e.ForEach("section", func(i int, colElement *colly.HTMLElement) {
			current := colElement.ChildText("h2")
			if current == "Current" {
				artist := colElement.ChildText("a")
				date := colElement.ChildText("p.card__date")
				paceExhibitions = append(paceExhibitions, Exhibition{
					Gallery:   "Alison Jacques",
					Location:  "22 Cork StreetLondon W1S 3NG",
					Artist:    artist,
					Title:     artist,
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
	SaveToExcel(paceExhibitions)
}
