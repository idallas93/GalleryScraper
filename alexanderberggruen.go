package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateAlexanderBerggruen(url string) {
	c := colly.NewCollector()
	var alexanderBerggruenExhibitions []Exhibition
	c.OnHTML("div.container", func(e *colly.HTMLElement) {
		e.ForEach("p", func(i int, colElement *colly.HTMLElement) {
			text := colElement.Text
			updatedText := strings.ReplaceAll(text, "\n", "%")
			splitText := strings.Split(updatedText, "%")
			if len(splitText) == 2 {
				artistAndTitle := splitText[0]
				date := splitText[1]
				splitArtistAndTitle := strings.Split(artistAndTitle, ":")
				if len(splitArtistAndTitle) == 2 {
					alexanderBerggruenExhibitions = append(alexanderBerggruenExhibitions, Exhibition{
						Gallery:   "Alexander Berggruen",
						Location:  "1018 Madison Ave Floor 3, New York, NY 10075",
						Artist:    splitArtistAndTitle[0],
						Title:     splitArtistAndTitle[1],
						StartDate: date,
						EndDate:   date,
						Notes:     "",
					})
				} else if len(splitArtistAndTitle) == 1 {
					alexanderBerggruenExhibitions = append(alexanderBerggruenExhibitions, Exhibition{
						Gallery:   "Alexander Berggruen",
						Location:  "1018 Madison Ave Floor 3, New York, NY 10075",
						Artist:    artistAndTitle,
						Title:     "Untitled",
						StartDate: date,
						EndDate:   date,
						Notes:     "",
					})
				}
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(alexanderBerggruenExhibitions)
}
