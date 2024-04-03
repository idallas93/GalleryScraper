package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateDavidZwirner(url string) {
	c := colly.NewCollector()

	var davidZwirnerExhibitions []Exhibition
	c.OnHTML("#listview", func(e *colly.HTMLElement) {
		e.ForEach("div.detail_left", func(i int, colElement *colly.HTMLElement) {
			locationText := colElement.ChildText("span.tabletCity")
			artistText := colElement.ChildText("div.item_title")
			dateText := colElement.ChildText("div.item_date_text")
			titleText := colElement.ChildText("div.item_sub_title")
			noteText := colElement.ChildText("")
			if dateText != "" {
				davidZwirnerExhibitions = append(davidZwirnerExhibitions, Exhibition{
					Gallery:   "David Zwirner",
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
	SaveToExcel(davidZwirnerExhibitions)
}
