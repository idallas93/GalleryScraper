package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateChateauShatto(url string) {
	c := colly.NewCollector()
	var chateauShattoExhibitions []Exhibition
	c.OnHTML("a:nth-of-type(1)", func(e *colly.HTMLElement) {
		date := e.ChildText("span:nth-of-type(1)")
		artistText := e.ChildText("span.artistName.curentArtist")
		title := e.ChildText("p.index-grid__text-subtitle")
		allText := e.ChildText("span")
		textWithoutDate := allText[len(date):]
		location := textWithoutDate[len(artistText):]
		if location != "" {
			chateauShattoExhibitions = append(chateauShattoExhibitions, Exhibition{
				Gallery:   "Chateau Shatto",
				Location:  location,
				Artist:    artistText,
				Title:     title,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(chateauShattoExhibitions)
}
