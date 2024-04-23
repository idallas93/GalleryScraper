package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updatePeresProjects(url string) {
	c := colly.NewCollector()
	var peresProjectsExhibitions []Exhibition
	c.OnHTML("#exhibitions-grid-forthcoming", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.subtitle")
			title := colElement.ChildText("h2")
			date := colElement.ChildText("span.date")
			location := colElement.ChildText("span.location")
			peresProjectsExhibitions = append(peresProjectsExhibitions, Exhibition{
				Gallery:   "Peres Projects",
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
	SaveToExcel(peresProjectsExhibitions)
}
