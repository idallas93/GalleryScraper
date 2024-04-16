package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateVictoriaMiro(url string) {
	c := colly.NewCollector()
	var victoriaMiroExhibitions []Exhibition
	c.OnHTML("#section-2", func(e *colly.HTMLElement) {
		e.ForEach("span.content", func(i int, colElement *colly.HTMLElement) {
			title := colElement.ChildText("h2.i")
			artistText := strings.ReplaceAll(colElement.ChildText("h2"), title, "")
			date := colElement.ChildText("span.top_subtitle")
			victoriaMiroExhibitions = append(victoriaMiroExhibitions, Exhibition{
				Gallery:   "Victoria Miro",
				Location:  "16 Wharf Rd, London N1 7RW, United Kingdom",
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
	SaveToExcel(victoriaMiroExhibitions)
}
