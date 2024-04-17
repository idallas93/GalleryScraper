package main

import (
	"log"
	"strconv"
	"github.com/gocolly/colly"
)

func updateCapitainPetzel(url string) {
	c := colly.NewCollector()
	var capitainPetzelExhibitions []Exhibition
	c.OnHTML("#main_content", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2")
			title := colElement.ChildText("span.subtitle")
			date := colElement.ChildText("span.date")
			var year string = ""
			if(len(date) > 4){
				year = date[len(date)-4:]
			}
			location := "Karl-Marx-Allee 45, 10178 Berlin, Germany"
			num, _ := strconv.Atoi(year)
			if year != "" && num >= 2024 {
				capitainPetzelExhibitions = append(capitainPetzelExhibitions, Exhibition{
					Gallery:   "Capitain Petzel",
					Location:  location,
					Artist:    artistText,
					Title:     title,
					StartDate: date,
					EndDate:   date,
					Notes:     "",
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(capitainPetzelExhibitions)
}
