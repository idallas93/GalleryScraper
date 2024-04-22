package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateGalerieBuchholz(url string) {
	c := colly.NewCollector()
	var galerieBuchholzExhibitions []Exhibition
	c.OnHTML("div._templatepart__pageexhibitions__list _exhibitionlist", func(e *colly.HTMLElement) {
		e.ForEach("div._templatepart__pageexhibitions__list__element _exhibitionlist__element", func(i int, colElement *colly.HTMLElement) {
			if i == 0 {
				artistText := colElement.ChildText("span")
				title := colElement.ChildText("p:nth-of-type(1)")
				location := colElement.ChildText("p:nth-of-type(2)")
				galerieBuchholzExhibitions = append(galerieBuchholzExhibitions, Exhibition{
					Gallery:   "Galerie Buchholz",
					Location:  location,
					Artist:    artistText,
					Title:     title,
					StartDate: "2024",
					EndDate:   "2024",
					Notes:     "",
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(galerieBuchholzExhibitions)
}
