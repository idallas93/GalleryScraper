package main

import (
	"log"

	"github.com/gocolly/colly"
)

func updateTimothyTaylor(url string) {
	c := colly.NewCollector()
	var timothyTaylorExhibitions []Exhibition
	c.OnHTML("#content", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("span:nth-of-type(2)")
			date := colElement.ChildText("span:nth-of-type(3)")
			location := colElement.ChildText("span:nth-of-type(4)")
			description := colElement.ChildText("div.description")
			if description != "" {
				timothyTaylorExhibitions = append(timothyTaylorExhibitions, Exhibition{
					Gallery:   "Timothy Taylor",
					Location:  location,
					Artist:    artistText,
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
	SaveToExcel(timothyTaylorExhibitions)
}
