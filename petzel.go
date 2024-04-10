package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updatePetzel(url string) {
	c := colly.NewCollector()
	var petzelExhibitions []Exhibition
	c.OnHTML("#large", func(e *colly.HTMLElement) {
		e.ForEach("div.headers", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			location := colElement.ChildText("h2.subtitle2")
			title := colElement.ChildText("h2")
			titleWithoutLocation := strings.ReplaceAll(title, location, "")
			date := colElement.ChildText("h3")
			petzelExhibitions = append(petzelExhibitions, Exhibition{
				Gallery:   "Petzel Gallery",
				Location:  location,
				Artist:    artistText,
				Title:     titleWithoutLocation,
				StartDate: date,
				EndDate:   date,
				Notes:     "",
			})
		})
	})
	c.OnHTML("#medium", func(f *colly.HTMLElement) {
		f.ForEach("div.headers", func(i int, fElement *colly.HTMLElement) {
			artistText := fElement.ChildText("h1")
			location := fElement.ChildText("h2.subtitle2")
			title := fElement.ChildText("h2")
			titleWithoutLocation := strings.ReplaceAll(title, location, "")
			date := fElement.ChildText("h3")
			petzelExhibitions = append(petzelExhibitions, Exhibition{
				Gallery:   "Petzel Gallery",
				Location:  location,
				Artist:    artistText,
				Title:     titleWithoutLocation,
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
	SaveToExcel(petzelExhibitions)
}
