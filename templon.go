package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateTemplon(url string) {
	c := colly.NewCollector()
	var templonExhibitions []Exhibition
	c.OnHTML("#a-venir", func(e *colly.HTMLElement) {
		e.ForEach("div.col-sm-6", func(i int, colElement *colly.HTMLElement) {
			locationAndDate := colElement.ChildText("p:nth-of-type(2)")
			locationAndDateSplit := strings.Split(locationAndDate, " | ")
			if(len(locationAndDateSplit) == 2){
				artistText := colElement.ChildText("p:nth-of-type(1)")
				title := colElement.ChildText("h3")
				date := locationAndDateSplit[0]
				location := locationAndDateSplit[1]
				templonExhibitions = append(templonExhibitions, Exhibition{
					Gallery:   "Templon",
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
	SaveToExcel(templonExhibitions)
}
