package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updatePippyHouldsworth(url string) {
	c := colly.NewCollector()
	var pippyHouldsworthExhibitions []Exhibition
	c.OnHTML("#content", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.ChildText("a:nth-of-type(1)")
			title := colElement.ChildText("span.h1_subtitle")
			artistText := strings.ReplaceAll(allText, title, "")
			artistText = strings.ReplaceAll(artistText, ":", "")
			date := colElement.ChildText("div.date")
			location := "6 Heddon St, London W1B 4BT, United Kingdom"
			pippyHouldsworthExhibitions = append(pippyHouldsworthExhibitions, Exhibition{
				Gallery:   "Pippy Houldsworth",
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
	SaveToExcel(pippyHouldsworthExhibitions)
}
