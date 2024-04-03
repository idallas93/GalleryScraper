package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateHauserAndWirthData(url string) {
	c := colly.NewCollector()

	var hauserAndWirthExhibitions []Exhibition
	c.OnHTML(".mb-13", func(e *colly.HTMLElement) {
		e.ForEach("div.col-span-full", func(i int, colElement *colly.HTMLElement) {
			dateText := colElement.ChildText("p.uppercase.text-black.font-medium")
			locationText := colElement.ChildText("p.mb-4.uppercase.text-black.font-medium")
			titleText := colElement.ChildText("p.mb-4.uppercase.text-black.font-medium")
			artistText := colElement.ChildText("a")
			notesText := colElement.ChildText("#text")	
			hauserAndWirthExhibitions = append(hauserAndWirthExhibitions, Exhibition{
				Gallery:   "Hauser & Wirth",
				Location:  locationText,
				Artist:    artistText,
				Title:     titleText,
				StartDate: dateText,
				EndDate:   dateText,
				Notes:     notesText,
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(hauserAndWirthExhibitions)
}
