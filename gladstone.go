package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGladstone(url string) {
	c := colly.NewCollector()
	var gladstoneExhibitions []Exhibition
	c.OnHTML("section.HomeSection.grid-container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("div.SmallTextCaps.small-TinyText.Bold")
			title := colElement.ChildText("div.SmallText.small-TinyText.StayBlack.Ital")
			date := colElement.ChildText("span")
			location := colElement.ChildText("div.smallText.small-TinyText.StayBlack")
			if(artistText != ""){
				gladstoneExhibitions = append(gladstoneExhibitions, Exhibition{
					Gallery:   "Gladstone",
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
	SaveToExcel(gladstoneExhibitions)
}
