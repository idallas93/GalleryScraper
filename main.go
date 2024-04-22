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
	lisson := "https://www.lissongallery.com/exhibitions"
	davidZwirner := "https://www.davidzwirner.com/exhibitions?view=upcoming"
	threeZeroThree := "https://www.303gallery.com/gallery-exhibitions/upcoming"
	pace := "https://www.pacegallery.com/exhibitions/"
	lehmannMaupin := "https://www.lehmannmaupin.com/exhibitions/upcoming"
	matthewMarksCurrent := "https://www.matthewmarks.com/exhibitions/current"
	matthewMarksUpcoming := "https://www.matthewmarks.com/exhibitions/current"
	almineRech := "https://www.alminerech.com/exhibitions/"
	karma := "https://karmakarma.org/exhibitions/"
	blum := "https://www.blum-gallery.com/exhibitions/categories/current"
	davidKordansky := "https://www.davidkordanskygallery.com/exhibitions"
	gladstone := "https://www.gladstonegallery.com/"
	maxHetzler := "https://www.maxhetzler.com/exhibitions"
	perrotin := "https://www.perrotin.com/exhibitions/upcoming"
	petzel := "https://www.petzel.com/exhibitions"
	skarstedt := "https://www.skarstedt.com/exhibitions"
	spruethMagers := "https://spruethmagers.com/exhibitions/upcoming"
	xavierHufkens := "https://www.xavierhufkens.com/exhibitions"
	francoisGhebaly := "https://ghebaly.com/"
	levyGoryDayan := "https://www.levygorvydayan.com/exhibitions"
	pilarCorrias := "https://www.pilarcorrias.com/exhibitions/"
	sadieColes := "https://www.sadiecoles.com/exhibitions/current/"
	luhringAugustine := "https://www.luhringaugustine.com/exhibitions/upcoming"
	victoriaMiro := "https://www.victoria-miro.com/exhibitions/"
	kurimanzutto := "https://www.kurimanzutto.com/exhibitions"
	kaufmannRepetto := "https://kaufmannrepetto.com/exhibitions-current/"
	barbati := "https://www.barbatigallery.com/"
	antonKern := "https://www.antonkerngallery.com/exhibitions/"
	allisonJacques := "https://alisonjacques.com/exhibitions"
	anatEbgi := "https://anatebgi.com/exhibitions/"
	vitoSchnabel := "https://www.vitoschnabel.com/exhibitions"
	bartolami := "https://www.bortolamigallery.com/exhibitions"
	cahiersDart := "https://www.cahiersdart.com/en/gallery/"
	canada := "https://www.canadanewyork.com/"
	capitainPetzel := "https://www.capitainpetzel.de/exhibitions/"
	caseyKaplan := "https://caseykaplangallery.com/exhibitions/"
	chantalCrousel := "https://www.crousel.com/en/exhibitions/"
	charlesMoffett := "https://charlesmoffett.com/exhibitions/"
	chateauShatto := "https://chateaushatto.com/"
	chrisSharpGallery := "https://www.chrissharpgallery.com/upcoming"
	clearing := "https://www.c-l-e-a-r-i-n-g.com/home-2/"
	craigStarr := "http://www.craigstarr.com/exhibitions"
	evaPresenhuber := "https://www.presenhuber.com/exhibitions"
	galerieBuchholz := "https://www.galeriebuchholz.de/exhibitions"
	giselaCapitain := "https://www.galeriecapitain.de/archive/?exhibition=upcoming&location=galerie-gisela-capitain"
	greeneNaftali := "https://greenenaftaligallery.com/exhibitions"
	thomasDaneGallery := "https://www.thomasdanegallery.com/exhibitions/"
	timothyTaylor := "https://www.timothytaylor.com/"
	utaArtistSpace := "https://utaartistspace.com/exhibition/"
	vsf := "https://www.vsf.la/exhibitions/"
	vielmetter := "https://vielmetter.com/exhibitions/"
	mendesWood := "https://mendeswooddm.com/exhibitions/"
	alexanderBerggruen := "https://alexanderberggruen.com/exhibitions/"
	hauserAndWorth := "https://www.hauserwirth.com/hauser-wirth-exhibitions/?date=forthcoming"
	grimm := "https://grimmgallery.com/exhibitions/"


	galleryUrls := []string{
		gCurrent, gUpcoming, whiteCube, lisson, davidZwirner,
		threeZeroThree, pace,
		lehmannMaupin, matthewMarksCurrent, matthewMarksUpcoming,
		almineRech, karma, blum, davidKordansky, gladstone, maxHetzler,
		perrotin, petzel, skarstedt, spruethMagers, xavierHufkens, francoisGhebaly,
		levyGoryDayan, pilarCorrias, sadieColes, luhringAugustine, victoriaMiro,
		kurimanzutto, kaufmannRepetto, barbati, antonKern, allisonJacques, anatEbgi,
		vitoSchnabel, bartolami, cahiersDart, canada, capitainPetzel, caseyKaplan,
		chantalCrousel, charlesMoffett, chateauShatto, chrisSharpGallery, clearing,
		craigStarr, evaPresenhuber, galerieBuchholz, giselaCapitain, greeneNaftali,
		thomasDaneGallery, timothyTaylor, utaArtistSpace, vsf, vielmetter,
		mendesWood, alexanderBerggruen, hauserAndWorth, grimm,
	}
	galleryFunctions := []func(url string){
		updateGagosianData, updateGagosianData, updateWhiteCubeData,
		updateLissonData, updateDavidZwirner,
		updateThreeZeroThreeData, updatePace,
		updateLehmannMaupin, updateMatthewMarks, updateMatthewMarks, updateAlmineRech,
		updateKarma, updateBlum, updateDavidKordansky, updateGladstone, updateMaxHetzler,
		updatePerrotin, updatePetzel, updateSkarstedt, updateSpruethMagers, updateXavierHufkens,
		updateFrancoisGhebaly, updateLevyGorvyDayan, updatePilarCorrias,
		updateSadieColes, updateLuhringAugustine, updateVictoriaMiro, updateKurimanzutto, 
		updateKaufmannRepetto, updateBarbati, updateAntonKern, updateAlisonJacques, updateAnatEbgi,
		updateVitoSchnabel, updateBartolami, updateCahiersDart, updateCanada, updateCapitainPetzel,
		updateCaseyKaplan, updateChantalCrousel, updateCharlieMoffett, updateChateauShatto, updateChrisSharpGallery,
		updateClearing, updateCraigStarr, updateEvaPresenhuber, updateGalerieBuchholz, updateGiselaCapitain, updateGreeneNaftali,
		updateThomasDaneGallery, updateTimothyTaylor, updateUtaArtistSpace, updateVsf, updateVielMetter,
		updateMendesWoodData, updateAlexanderBerggruen, updateHauserAndWirthData, updateGrimm,
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
	f, err := excelize.OpenFile("galleryData.xlsx")
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
		if err := f.SaveAs("galleryData.xlsx"); err != nil {
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
		if err := f.SaveAs("galleryData.xlsx"); err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
