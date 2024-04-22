package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateVielMetter(url string) {
	c := colly.NewCollector()
	var vielmetterExhibitions []Exhibition
	c.OnHTML("#exhibitions-current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			text  := colElement.ChildText("p:nth-of-type(2)")
			splitText := strings.Split(text, ": ")
			if(len(splitText) == 2){
				artist := splitText[0]
				title := splitText[1]
				date := colElement.ChildText("p:nth-of-type(3)")
				location := "1700 S Santa Fe Ave #101"
				vielmetterExhibitions = append(vielmetterExhibitions, Exhibition{
					Gallery:   "Vielmetter",
					Location:  location,
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
	SaveToExcel(vielmetterExhibitions)
}
