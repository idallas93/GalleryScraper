package main

import (
	"log"
	"regexp"
	"strings"
	"github.com/gocolly/colly"
)

func updateWhiteCubeData(url string) {
	c := colly.NewCollector()
	var whiteCubeExhibitions []Exhibition
	c.OnHTML("#main", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		allText := e.ChildText("p")
		space := regexp.MustCompile(`\s{2,}`)
		allText = space.ReplaceAllString(allText, ";")
		splitText := strings.Split(allText, "Gallery Exhibition;")
		for _, exhibition := range splitText {
			splitExhibitionText := strings.Split(exhibition, ";")
			 if len(splitExhibitionText) == 5 {
				artist := splitExhibitionText[0]
				title := splitExhibitionText[1]
				dates := splitExhibitionText[2]
				location := splitExhibitionText[3]
				whiteCubeExhibitions = append(whiteCubeExhibitions, Exhibition{
					Gallery:   "White Cube",
					Location:  location,
					Artist:    artist,
					Title:     title,
					StartDate: dates,
					EndDate:   dates,
					Notes:     "",
				})
			} else if len(splitExhibitionText) == 4 {
				artist := splitExhibitionText[0]
				dates := splitExhibitionText[1]
				location := splitExhibitionText[2]
				whiteCubeExhibitions = append(whiteCubeExhibitions, Exhibition{
					Gallery:   "White Cube",
					Location:  location,
					Artist:    artist,
					Title:     "Untitled",
					StartDate: dates,
					EndDate:   dates,
					Notes:     "",
				})
			} else if len(splitExhibitionText) == 3 {
				artist := splitExhibitionText[0]
				location := splitExhibitionText[1]
				whiteCubeExhibitions = append(whiteCubeExhibitions, Exhibition{
					Gallery:   "White Cube",
					Location:  location,
					Artist:    artist,
					Title:     "Untitled",
					StartDate: "",
					EndDate:   "",
					Notes:     "",
				})
			} else if len(splitExhibitionText) < 3 && len(splitExhibitionText) > 1 {
				artist := splitExhibitionText[0]
				location := splitExhibitionText[1]
				whiteCubeExhibitions = append(whiteCubeExhibitions, Exhibition{
					Gallery:   "White Cube",
					Location:  location,
					Artist:    artist,
					Title:     "Untitled",
					StartDate: "",
					EndDate:   "",
					Notes:     "",
				})
			}
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(whiteCubeExhibitions)
}
