package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updatePaulaCooperGallery(url string) {
	c := colly.NewCollector()
	var paulaCooperExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("h3")
			location := colElement.ChildText("h2.subtitle2")
			if colElement.ChildText("h2:nth-of-type(2)") != "" {
				title = strings.ReplaceAll(title, location, "")
			}
			paulaCooperExhibitions = append(paulaCooperExhibitions, Exhibition{
				Gallery:   "Paula Cooper Gallery",
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
	SaveToExcel(paulaCooperExhibitions)
}
