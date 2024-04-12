package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateLevyGorvyDayan(url string) {
	c := colly.NewCollector()
	var levyGorvyDayanExhibitions []Exhibition
	c.OnHTML("#current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.Text
			title := colElement.ChildText("li.italic")
			location := strings.ReplaceAll(allText, title, "")
			levyGorvyDayanExhibitions = append(levyGorvyDayanExhibitions, Exhibition{
				Gallery:   "Levy Gorvy Dayan",
				Location:  location,
				Artist:    title,
				Title:     title,
				StartDate: location,
				EndDate:   location,
				Notes:     "",
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(levyGorvyDayanExhibitions)
}
