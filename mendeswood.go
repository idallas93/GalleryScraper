package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateMendesWoodData(url string) {
	c := colly.NewCollector()

	var mendesWoodExhibitions []Exhibition
	c.OnHTML("#main_content", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			locationText := colElement.ChildText(".location")
			artistText := colElement.ChildText(".subtitle")
			dateText := colElement.ChildText(".date")
			titleText := colElement.ChildText("h2")
			noteText := colElement.ChildText(".description.pose")
			if dateText != "" {
				mendesWoodExhibitions = append(mendesWoodExhibitions, Exhibition{
					Gallery:   "Mendes Wood DM",
					Location:  locationText,
					Artist:    artistText,
					Title:     titleText,
					StartDate: dateText,
					EndDate:   dateText,
					Notes:     noteText,
				})
			}

		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(mendesWoodExhibitions)
}
