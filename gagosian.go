package main

import (
	// "fmt"
	"log"
	"regexp"
	"strings"
	"github.com/gocolly/colly"
)

func updateGagosianData(url string) {
	c := colly.NewCollector()
	var gagosianExhibitions []Exhibition
	c.OnHTML(".ex-item", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		allText := e.ChildText("a")
		space := regexp.MustCompile(`\s{2,}`)
		allText = space.ReplaceAllString(allText, ";")
		splitText := strings.Split(allText, ";")
		if len(splitText) > 4 {
			notes := splitText[0]
			artist := splitText[1]
			title := splitText[2]
			dates := splitText[3]
			location := splitText[4]
			gagosianExhibitions = append(gagosianExhibitions, Exhibition{
				Gallery:   "Gagosian",
				Location:  location,
				Artist:    artist,
				Title:     title,
				StartDate: dates,
				EndDate:   dates,
				Notes:     notes,
			})
		} else if len(splitText) > 3 {
			artist := splitText[0]
			title := splitText[1]
			dates := splitText[2]
			location := splitText[3]
			gagosianExhibitions = append(gagosianExhibitions, Exhibition{
				Gallery:   "Gagosian",
				Location:  location,
				Artist:    artist,
				Title:     title,
				StartDate: dates,
				EndDate:   dates,
				Notes:     "",
			})
		} else if len(splitText) > 2 {
			artist := splitText[0]
			dates := splitText[1]
			location := splitText[2]
			gagosianExhibitions = append(gagosianExhibitions, Exhibition{
				Gallery:   "Gagosian",
				Location:  location,
				Artist:    artist,
				Title:     "Untitled",
				StartDate: dates,
				EndDate:   dates,
				Notes:     "",
			})
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(gagosianExhibitions)
}
