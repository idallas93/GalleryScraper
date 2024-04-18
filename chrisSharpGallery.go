package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateChrisSharpGallery(url string) {
	c := colly.NewCollector()
	var chrisSharpGalleryExhibitions []Exhibition
	c.OnHTML("article.sections", func(e *colly.HTMLElement) {
		e.ForEach("div.content", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("p:nth-of-type(1)")
			title := ""
			date := colElement.ChildText("p:nth-of-type(2)")
			location :=  "4650 W Washington BlvdLos Angeles, CA 90016"
			chrisSharpGalleryExhibitions = append(chrisSharpGalleryExhibitions, Exhibition{
				Gallery:   "Chris Sharp Gallery",
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
	SaveToExcel(chrisSharpGalleryExhibitions)
}
