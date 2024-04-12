package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateSpruethMagers(url string) {
	c := colly.NewCollector()
	var spruethMagersExhibitions []Exhibition
	c.OnHTML("#main", func(e *colly.HTMLElement) {
		e.ForEach("article", func(i int, colElement *colly.HTMLElement) {
			location := colElement.ChildText("p.location-p")
			artist := colElement.ChildText("b")
			title := colElement.ChildText("i")
			allText := colElement.ChildText("p")
			spruethMagersExhibitions = append(spruethMagersExhibitions, Exhibition{
					Gallery:   "Sprueth Magers",
					Location:  location,
					Artist:    artist,
					Title:     title,
					StartDate: allText,
					EndDate:   "",
					Notes:     "",
				})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(spruethMagersExhibitions)
}
