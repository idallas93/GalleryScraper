package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGladstone(url string) {
	c := colly.NewCollector()
	var gladstoneExhibitions []Exhibition
	c.OnHTML("#main_content", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("div.SmallTextCaps.Bold")
			title := colElement.ChildText("div.SmallText.StayBlack.Ital")
			date := colElement.ChildText("span")
			location := colElement.ChildText("div.SmallText.StayBlack")
			gladstoneExhibitions = append(gladstoneExhibitions, Exhibition{
				Gallery:   "Gladstone",
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
	SaveToExcel(gladstoneExhibitions)
}
