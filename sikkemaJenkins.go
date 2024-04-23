package main

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func updateSikkemaJenkins(url string) {
	c := colly.NewCollector()
	var sikkemaJenkinsExhibitions []Exhibition
	c.OnHTML("#page", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.Text
			title := colElement.ChildText("em")
			if(title != ""){
				splitAllTextByTitle := strings.Split(allText, title)
				if len(splitAllTextByTitle) == 2{
					artist := splitAllTextByTitle[0]
					date := splitAllTextByTitle[1]
					
					sikkemaJenkinsExhibitions = append(sikkemaJenkinsExhibitions, Exhibition{
						Gallery:   "Sikkema Jenkins & Co.",
						Location:  "530 WEST 22ND STREET NEW YORK NY 10011",
						Artist:    artist,
						Title:     title,
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
	SaveToExcel(sikkemaJenkinsExhibitions)
}
