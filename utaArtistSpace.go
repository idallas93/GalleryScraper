package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateUtaArtistSpace(url string) {
	c := colly.NewCollector()
	var utaArtistSpaceExhibitions []Exhibition
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("div.exhibition-index-container.grid-x", func(i int, colElement *colly.HTMLElement) {
			if(i == 0){
				colElement.ForEach("a", func(j int, f *colly.HTMLElement) {
					artistText := f.ChildText("h2")
					title := f.ChildText("h1")
					date := f.ChildText("p:nth-of-type(1)")
					location := f.ChildText("p:nth-of-type(2)")
					utaArtistSpaceExhibitions = append(utaArtistSpaceExhibitions, Exhibition{
						Gallery:   "UTA Artist Space",
						Location:  location,
						Artist:    artistText,
						Title:     title,
						StartDate: date,
						EndDate:   date,
						Notes:     "",
					})
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(utaArtistSpaceExhibitions)
}
