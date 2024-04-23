package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateStephenFriedman(url string) {
	c := colly.NewCollector()
	var stephenFriedmanExhibitions []Exhibition
	c.OnHTML("panel_number_0.panel_index_3.has_panel_heading", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("span.heading_title")
			locationAndDate := colElement.ChildText("div.subheading.ani-in")
			splitLocationAndDate := strings.Split(locationAndDate, ", ")
			if len(splitLocationAndDate) == 3{
				location := splitLocationAndDate[0]
				date := splitLocationAndDate[1]
				title := colElement.ChildText("p.index-grid__text-subtitle")
				stephenFriedmanExhibitions = append(stephenFriedmanExhibitions, Exhibition{
					Gallery:   "Stephen Friedman Gallery",
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
	SaveToExcel(stephenFriedmanExhibitions)
}
