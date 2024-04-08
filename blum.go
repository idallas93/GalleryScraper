package main

import (
	"log"
	"strings"
	"github.com/gocolly/colly"
)

func updateBlum(url string) {
	c := colly.NewCollector()
	var blumExhibitions []Exhibition
	c.OnHTML("#exhibitions_index_wrapper", func(e *colly.HTMLElement) {
		e.ForEach("a", func(i int, colElement *colly.HTMLElement) {
			artistText := colElement.ChildText("h2.artist_name")
			title := colElement.ChildText("div.italic")
			// dateAndLocation := colElement.ChildText("h2.div")
			galleryTextWithoutTitle := strings.ReplaceAll(colElement.Text, title, "")
			galleryTextWithoutTitleAndArtistText := strings.ReplaceAll(galleryTextWithoutTitle, artistText, "")
			containsTokyo := strings.Contains(galleryTextWithoutTitleAndArtistText, "Tokyo")
			containsNewYork := strings.Contains(galleryTextWithoutTitleAndArtistText, "New York")
			containsLA := strings.Contains(galleryTextWithoutTitleAndArtistText, "Los Angeles")
			datesWithoutTitleArtistTextAndTokyo := strings.ReplaceAll(galleryTextWithoutTitleAndArtistText, "Tokyo", "")
			datesWithoutTitleArtistTextAndLA := strings.ReplaceAll(galleryTextWithoutTitleAndArtistText, "Los Angeles", "")
			datesWithoutTitleArtistTextAndNewYork := strings.ReplaceAll(galleryTextWithoutTitleAndArtistText, "New York", "")
			if containsTokyo {
				blumExhibitions = append(blumExhibitions, Exhibition{
					Gallery:   "Blum",
					Location:  "Tokyo",
					Artist:    artistText,
					Title:     title,
					StartDate: datesWithoutTitleArtistTextAndTokyo,
					EndDate:   datesWithoutTitleArtistTextAndTokyo,
					Notes:     "",
				})
			} else if containsLA {
				blumExhibitions = append(blumExhibitions, Exhibition{
					Gallery:   "Blum",
					Location:  "Los Angeles",
					Artist:    artistText,
					Title:     title,
					StartDate: datesWithoutTitleArtistTextAndLA,
					EndDate:   datesWithoutTitleArtistTextAndLA,
					Notes:     "",
				})
			} else if containsNewYork {
				blumExhibitions = append(blumExhibitions, Exhibition{
					Gallery:   "Blum",
					Location:  "New York",
					Artist:    artistText,
					Title:     title,
					StartDate: datesWithoutTitleArtistTextAndNewYork,
					EndDate:   datesWithoutTitleArtistTextAndNewYork,
					Notes:     "",
				})
			}
		})
	})
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
	SaveToExcel(blumExhibitions)
}
