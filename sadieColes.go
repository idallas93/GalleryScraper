package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateSadieColes(url string) {
	c := colly.NewCollector()
	var sadieColesExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("span.date")
			location := colElement.ChildText("span.location")
			allText := strings.ReplaceAll(colElement.ChildText("h2"), title, "")
			allTextWithoutDate := strings.ReplaceAll(allText, date, "")
			artistText := strings.ReplaceAll(allTextWithoutDate, location, "")
			sadieColesExhibitions = append(sadieColesExhibitions, Exhibition{
				Gallery:   "Sadie Coles",
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
	SaveToExcel(sadieColesExhibitions)
}
