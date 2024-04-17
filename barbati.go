package main

import (
	"log"
	"github.com/gocolly/colly"
)

func updateBarbati(url string) {
	c := colly.NewCollector()
	var barbatiExhibitions []Exhibition
	c.OnHTML("#tm-main", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h1")
			title := colElement.ChildText("div.el-meta.uk-h3.uk-margin-top.uk-margin-remove-bottom")
			date := colElement.ChildText("div.el-content.uk-panel.uk-margin-top")
			location := "Los Angeles, CA"
			barbatiExhibitions = append(barbatiExhibitions, Exhibition{
				Gallery:   "Barbati",
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
	SaveToExcel(barbatiExhibitions)
}
