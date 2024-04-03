package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateThreeZeroThreeData(url string) {
	c := colly.NewCollector()
	var threeZeroThreeExhibitions []Exhibition
	c.OnHTML("#exhibitions-container", func(e *colly.HTMLElement) {
		e.ForEach("div.headers", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			dateText := colElement.ChildText("h3")
			titleText := colElement.ChildText("h2")
			noteText := colElement.ChildText("")
				threeZeroThreeExhibitions = append(threeZeroThreeExhibitions, Exhibition{
					Gallery:   "303",
					Location:  "555 W 21 Street New York",
					Artist:    artistText,
					Title:     titleText,
					StartDate: dateText,
					EndDate:   dateText,
					Notes:     noteText,
				})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(threeZeroThreeExhibitions)
}
