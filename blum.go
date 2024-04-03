package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func updateBlum(url string) {
	c := colly.NewCollector()
	var blumExhibitions []Exhibition
	c.OnHTML("#exhibitions_index_wrapper", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2.artist_name")
			title := colElement.ChildText("div.italic")
			// dateAndLocation := colElement.ChildText("h2.div")
			fmt.Println(colElement.Text)
			location := colElement.ChildText("p.index-grid__text-location")
			blumExhibitions = append(blumExhibitions, Exhibition{
				Gallery:   "Pace Gallery",
				Location:  location,
				Artist:    artistText,
				Title:     title,
				StartDate: "",
				EndDate:   "",
				Notes:     "",
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(blumExhibitions)
}
