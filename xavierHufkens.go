package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateXavierHufkens(url string) {
	c := colly.NewCollector()
	var xavierHufkens []Exhibition
	c.OnHTML("div.o-wrapper", func(e *colly.HTMLElement) {
		e.ForEach("#exhibitions-list", func(i int, f *colly.HTMLElement) {
			if i == 0 || i == 1 {
				f.ForEach("div.c-card-exhibition__info", func(j int, g *colly.HTMLElement) {
					spanText := g.ChildText("span")
					title := g.ChildText("span.c-card-exhibition__title")
					date := strings.ReplaceAll(g.ChildText("time"), title, "")
					locationWithoutDate := strings.ReplaceAll(g.ChildText("div.c-card-exhibition__data.u-body-sans-small"), date, "")
					location := strings.ReplaceAll(locationWithoutDate, "Until  ", "")
					artist := strings.ReplaceAll(spanText, title, "")
					artistWithoutDate := strings.ReplaceAll(artist, date, "")
					artistText := strings.ReplaceAll(artistWithoutDate, "Until  ","")
					xavierHufkens = append(xavierHufkens, Exhibition{
						Gallery:   "Xavier Hufkens",
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
	SaveToExcel(xavierHufkens)
}
