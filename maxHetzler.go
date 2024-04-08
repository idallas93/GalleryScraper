package main

import (
	"log"

	"github.com/gocolly/colly"
)

func updateMaxHetzler(url string) {
	c := colly.NewCollector()
	var maxHetzler []Exhibition
	c.OnHTML("div.row.current-exhibitions", func(e *colly.HTMLElement) {
		e.ForEach("div.col-sm-12.info", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h3")
			title := colElement.ChildText("p")
			date := colElement.ChildText("div.date")
			location := colElement.ChildText("div.loc")
			maxHetzler = append(maxHetzler, Exhibition{
				Gallery:   "Max Hetzler",
				Location:  location,
				Artist:    artistText,
				Title:     title,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		})
	})
	c.OnHTML("div.row.upcoming-exhibitions", func(f *colly.HTMLElement) {
		f.ForEach("div.col-sm-12.info", func(i int, fElement *colly.HTMLElement) {
			upcomingArtistText := fElement.ChildText("h3")
			upcomingTitle := fElement.ChildText("p")
			upcomingDate := fElement.ChildText("div.date")
			location := fElement.ChildText("div.loc")
			maxHetzler = append(maxHetzler, Exhibition{
				Gallery:   "Max Hetzler",
				Location:  location,
				Artist:    upcomingArtistText,
				Title:     upcomingTitle,
				StartDate: upcomingDate,
				EndDate:   upcomingDate,
				Notes:     "",
			})
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(maxHetzler)
}
