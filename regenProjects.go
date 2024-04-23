package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateRegenProjects(url string) {
	c := colly.NewCollector()
	var regenProjectsExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("h3")
			location := "6750 Santa Monica Boulevard Los Angeles, CA 90038"
			regenProjectsExhibitions = append(regenProjectsExhibitions, Exhibition{
				Gallery:   "Regen Projects",
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
	SaveToExcel(regenProjectsExhibitions)
}
