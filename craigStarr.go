package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateCraigStarr(url string) {
	c := colly.NewCollector()
	var craigStarrExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			titleAndText :=colElement.ChildText("h1")
			artistText := strings.Split(titleAndText, ":")[0]
			title := strings.Split(titleAndText, ":")[1]
			date := colElement.ChildText("h3")
			location := "5 EAST 73RD STREET, NEW YORK, NY 10021"
			craigStarrExhibitions = append(craigStarrExhibitions, Exhibition{
				Gallery:   "Craig Starr",
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
	SaveToExcel(craigStarrExhibitions)
}
