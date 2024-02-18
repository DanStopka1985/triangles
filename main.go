package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	tsApp        fyne.App
	canvasWindow fyne.Window
	btnsWindow   fyne.Window
	canvasCont   *fyne.Container

	ticker *time.Ticker
)

func initialization() {
	tsApp = app.New()
	rand.Seed(time.Now().UnixNano())
	ticker = time.NewTicker(evolutionSpeed)
	canvasWindow = tsApp.NewWindow("triangles")
	canvasWindow.Resize(fyne.NewSize(side+100, side+100))
	btnsWindow = tsApp.NewWindow("buttons")
	btnsWindow.Resize(fyne.NewSize(400, 400))
	canvasCont = container.NewWithoutLayout()
}

func main() {
	initialization()

	maxPopulationEntry := widget.NewEntry()
	maxPopulationEntry.SetText(strconv.Itoa(maxPopulation))
	maxPopulationEntry.OnChanged = func(s string) {
		log.Println("changed")
	}

	maxPopulationEntry.OnSubmitted = func(s string) {
		newVal, err := strconv.Atoi(s)
		if err != nil {
			// ... handle error
			panic(err)
		}
		maxPopulation = newVal
		log.Println("submitted")
	}

	startTsCntEntry := widget.NewEntry()
	startTsCntEntry.SetText(strconv.Itoa(startTsCnt))
	startTsCntEntry.OnChanged = func(s string) {
		log.Println("changed")
	}

	startTsCntEntry.OnSubmitted = func(s string) {
		newVal, err := strconv.Atoi(s)
		if err != nil {
			startTsCnt = startTsCntDefault
			startTsCntEntry.SetText(strconv.Itoa(startTsCnt))

		}
		startTsCnt = newVal
		log.Println("submitted")
	}

	form := widget.NewForm(widget.NewFormItem("Population", maxPopulationEntry),
		widget.NewFormItem("start triangles count (numeric)", startTsCntEntry))

	btnsWindow.SetContent(
		container.NewVBox(
			form, refreshButton(),
			genRandomButton(),
			startEvolutionButton(),
			stopEvolutionButton(),
		),
	)

	canvasWindow.SetContent(canvasCont)

	canvasWindow.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})

	canvasWindow.Show()
	btnsWindow.ShowAndRun()
}
