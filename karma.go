package main

import (
	"log"
	"regexp"
	"strings"
	"github.com/gocolly/colly"
)

func updateKarma(url string) {
	c := colly.NewCollector()
	var karmaExhibitions []Exhibition
	c.OnHTML("#section-current", func(e *colly.HTMLElement) {
		e.ForEach("div.link-wrap", func(i int, colElement *colly.HTMLElement) {
			allText := colElement.Text
			updatedText := strings.ReplaceAll(allText, "\n", "%")
			space := regexp.MustCompile(`\s{2,}`)
			allTextWithoutExtraSpaces := space.ReplaceAllString(updatedText, "")
			splitText := strings.Split(allTextWithoutExtraSpaces, "%")
			textWithoutEmptyVals := delete_empty(splitText)
			if len(textWithoutEmptyVals) == 4 {
			karmaExhibitions = append(karmaExhibitions, Exhibition{
					Gallery:   "Karma",
					Location:  textWithoutEmptyVals[3],
					Artist:    textWithoutEmptyVals[0],
					Title:     textWithoutEmptyVals[1],
					StartDate: textWithoutEmptyVals[2],
					EndDate:   textWithoutEmptyVals[2],
					Notes:     "",
				})
			} else if len(textWithoutEmptyVals) == 3 {
				karmaExhibitions = append(karmaExhibitions, Exhibition{
					Gallery:   "Karma",
					Location:  textWithoutEmptyVals[2],
					Artist:    textWithoutEmptyVals[0],
					Title:     "Untitled",
					StartDate: textWithoutEmptyVals[1],
					EndDate:   textWithoutEmptyVals[1],
					Notes:     "",
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(karmaExhibitions)
}
