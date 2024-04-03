package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateMatthewMarks(url string) {
	c := colly.NewCollector()
	var matthewMarksExhibitions []Exhibition
	c.OnHTML("#exhibitions_current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			paragraphText := colElement.ChildText("p")
			title := colElement.ChildText("em")
			artist := strings.ReplaceAll(paragraphText, title, "")
			date := colElement.ChildText("h2")
			locationGroup := colElement.ChildText("h3.location.gray.showCityStreet")
			if locationGroup != "" {
				matthewMarksExhibitions = append(matthewMarksExhibitions, Exhibition{
					Gallery:   "Matthew Marks",
					Location:  locationGroup,
					Artist:    artist,
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
	SaveToExcel(matthewMarksExhibitions)
}
