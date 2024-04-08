package main

import (
	"fmt"
	"strconv"
	"github.com/360EntSecGroup-Skylar/excelize"
)

type Exhibition struct {
	Gallery   string
	Location  string
	Artist    string
	Title     string
	StartDate string
	EndDate   string
	Notes     string
}

func main() {
	gCurrent := "https://gagosian.com/exhibitions/"
	gUpcoming := "https://gagosian.com/exhibitions/upcoming/"
	whiteCube := "https://www.whitecube.com/exhibitions/upcoming"
	hauserAndWorth := "https://www.hauserwirth.com/hauser-wirth-exhibitions/?date=forthcoming"
	lisson := "https://www.lissongallery.com/exhibitions"
	davidZwirner := "https://www.davidzwirner.com/exhibitions?view=upcoming"
	threeZeroThree := "https://www.303gallery.com/gallery-exhibitions/upcoming"
	alexanderBerggruen := "https://alexanderberggruen.com/exhibitions/"
	pace := "https://www.pacegallery.com/exhibitions/"
	lehmannMaupin := "https://www.lehmannmaupin.com/exhibitions/upcoming"
	matthewMarksCurrent := "https://www.matthewmarks.com/exhibitions/current"
	matthewMarksUpcoming := "https://www.matthewmarks.com/exhibitions/current"
	almineRech := "https://www.alminerech.com/exhibitions/"
	karma := "https://karmakarma.org/exhibitions/"
	blum := "https://www.blum-gallery.com/exhibitions/categories/current"
	// mendesWood := "https://mendeswooddm.com/exhibitions/"
	galleryUrls := []string{
		gCurrent, gUpcoming, whiteCube,
		hauserAndWorth, lisson, davidZwirner,
		threeZeroThree, alexanderBerggruen, pace,
		lehmannMaupin, matthewMarksCurrent, matthewMarksUpcoming,
		almineRech, karma, blum, 
	}
	galleryFunctions := []func(url string){
		updateGagosianData, updateGagosianData, updateWhiteCubeData,
		updateHauserAndWirthData, updateLissonData, updateDavidZwirner,
		updateThreeZeroThreeData, updateAlexanderBerggruen, updatePace,
		updateLehmannMaupin, updateMatthewMarks, updateMatthewMarks, updateAlmineRech,
		updateKarma, updateBlum,
	}
	if len(galleryUrls) == len(galleryFunctions) {
		for i := 0; i < len(galleryUrls); i++ {
			for j := 0; j < len(galleryFunctions); j++ {
				if i == j {
					galleryFunctions[j](galleryUrls[i])
				}
			}
		}
	} else {
		fmt.Println("lengths of gallery urls and functions urls are different")
	}
}

func SaveToExcel(exhibition []Exhibition) error {
	f, err := excelize.OpenFile("inventory.xlsx")
	if err != nil {
		f = excelize.NewFile()
		f.SetCellValue("Sheet1", "A1", "Gallery")
		f.SetCellValue("Sheet1", "B1", "Location")
		f.SetCellValue("Sheet1", "C1", "Artist")
		f.SetCellValue("Sheet1", "D1", "Title")
		f.SetCellValue("Sheet1", "E1", "StartDate")
		f.SetCellValue("Sheet1", "F1", "EndDate")

		for i := 0; i < len(exhibition); i++ {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), exhibition[i].Gallery)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), exhibition[i].Location)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), exhibition[i].Artist)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), exhibition[i].Title)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), exhibition[i].StartDate)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), exhibition[i].EndDate)
		}
		if err := f.SaveAs("inventory.xlsx"); err != nil {
			fmt.Println(err)
		}
	} else {
		rows := f.GetRows("Sheet1")
		for i := 0; i < len(exhibition); i++ {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(len(rows)+1+i), exhibition[i].Gallery)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(len(rows)+1+i), exhibition[i].Location)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(len(rows)+1+i), exhibition[i].Artist)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(len(rows)+1+i), exhibition[i].Title)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(len(rows)+1+i), exhibition[i].StartDate)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(len(rows)+1+i), exhibition[i].EndDate)
		}
		if err := f.SaveAs("inventory.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
