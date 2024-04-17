package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateAnatEbgi(url string) {
	c := colly.NewCollector()
	var anatEbgiExhibitions []Exhibition
	c.OnHTML("#exhibitions-current", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.ChildText("p.font-heading.pt-3.uppercase")
			var artist string
			var title string
			splitAllText := strings.Split(allText, ":")
			if len(splitAllText) == 2 {
				artist = splitAllText[0]
				title = splitAllText[1]
				dateAndLocation := colElement.ChildText("p.text-sm.font-heading.uppercase")
				splitDateAndLocation := strings.Split(dateAndLocation, ",")
				date := splitDateAndLocation[0]
				location := splitDateAndLocation[1]
				anatEbgiExhibitions = append(anatEbgiExhibitions, Exhibition{
					Gallery:   "Anat Ebgi",
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
	SaveToExcel(anatEbgiExhibitions)
}
