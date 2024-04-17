package main

import (
	"log"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

func updateVitoSchnabel(url string) {
	c := colly.NewCollector()
	var vitoSchnabelExhibitions []Exhibition
	c.OnHTML("#exhibitions-container", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			yearSplit := strings.Split(colElement.ChildText("h3"), ",")
			if len(yearSplit) == 2 {
				year := strings.ReplaceAll(yearSplit[1], " ", "")
				num, _ := strconv.Atoi(year)
				if num >= 2024 {
					artist := colElement.ChildText("h1")
					location := colElement.ChildText("h2.subtitle2")
					title := strings.ReplaceAll(colElement.ChildText("h2"), location, "")
					date := colElement.ChildText("h3")
					if artist != "" {
						vitoSchnabelExhibitions = append(vitoSchnabelExhibitions, Exhibition{
							Gallery:   "Vito Schnabel",
							Location:  location,
							Artist:    artist,
							Title:     title,
							StartDate: date,
							EndDate:   date,
							Notes:     "",
						})
					}
				}
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(vitoSchnabelExhibitions)
}
