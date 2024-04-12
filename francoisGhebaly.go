package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateFrancoisGhebaly(url string) {
	c := colly.NewCollector()
	var francoisGhebalyExhibitions []Exhibition
	c.OnHTML("#content-holder", func(e *colly.HTMLElement) {
		e.ForEach("p", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.semibold")
			title := colElement.ChildText("span.custom_one_italic")
			allText := colElement.ChildText("span")
			dateAndLocationWithoutArist := strings.ReplaceAll(allText, artistText, "")
			dateAndLocation := strings.ReplaceAll(dateAndLocationWithoutArist, title, "")
			var date string
			var location string

			if strings.Contains(dateAndLocation, "391") {
				dateAndLocationSlice := strings.Split(dateAndLocation, "391")
				date = dateAndLocationSlice[0]
				location = "391" + dateAndLocationSlice[1]
			} else if strings.Contains(dateAndLocation, "2245") {
				dateAndLocationSlice := strings.Split(dateAndLocation, "2245")
				date = dateAndLocationSlice[0]
				location = "2245" + dateAndLocationSlice[1]
			}
			if location != "" {
				francoisGhebalyExhibitions = append(francoisGhebalyExhibitions, Exhibition{
					Gallery:   "Francois Ghebaly",
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
	SaveToExcel(francoisGhebalyExhibitions)
}
