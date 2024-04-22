package main

import (
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

func updateClearing(url string) {
	c := colly.NewCollector()
	var clearingExhibitions []Exhibition
	c.OnHTML("#contenair", func(e *colly.HTMLElement) {
		e.ForEach("section", func(i int, colElement *colly.HTMLElement) {
			date := colElement.ChildText("p:nth-of-type(2)")
			var year string
			if len(date) >= 4{
				year = date[len(date)-4:]
			}
			num, _ := strconv.Atoi(year)
			if year != "" && num >= 2024 {
				artistText := colElement.ChildText("p:nth-of-type(1)")
				title := ""
				location := colElement.ChildText("p:nth-of-type(3)")
				clearingExhibitions = append(clearingExhibitions, Exhibition{
					Gallery:   "Clearing",
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
	SaveToExcel(clearingExhibitions)
}
