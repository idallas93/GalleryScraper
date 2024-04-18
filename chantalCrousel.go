package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateChantalCrousel(url string) {
	c := colly.NewCollector()
	var chantalCrouselExhibitions []Exhibition
	c.OnHTML("section:nth-of-type(2)", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("h3")
			date := colElement.ChildText("p")
			location := "10 RUE CHARLOT, 75003 PARIS"
			chantalCrouselExhibitions = append(chantalCrouselExhibitions, Exhibition{
				Gallery:   "Chantal Crousel",
				Location:  location,
				Artist:    artistText,
				Title:     title,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(chantalCrouselExhibitions)
}
