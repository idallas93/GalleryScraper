package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateCharlieMoffett(url string) {
	c := colly.NewCollector()
	var charlieMoffettExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-current", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("span.date")
			location := "431 Washington Street New York, NY 10013"
			charlieMoffettExhibitions = append(charlieMoffettExhibitions, Exhibition{
				Gallery:   "Charlie Moffett",
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
	SaveToExcel(charlieMoffettExhibitions)
}
