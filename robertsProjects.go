package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateRobertsProjects(url string) {
	c := colly.NewCollector()
	var robertsProjectsExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("h3")
			location := "442 South La Brea Avenue, Los Angeles, CA 90036"
			robertsProjectsExhibitions = append(robertsProjectsExhibitions, Exhibition{
				Gallery:   "Roberts Projects",
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
	SaveToExcel(robertsProjectsExhibitions)
}
