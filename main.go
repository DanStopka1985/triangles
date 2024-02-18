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
	tsApp      fyne.App
	canvasWin  fyne.Window
	ctrlWin    fyne.Window
	canvasCont *fyne.Container
	inputForm  *widget.Form

	ticker *time.Ticker
)

func initialization() {
	tsApp = app.New()
	rand.Seed(time.Now().UnixNano())
	ticker = time.NewTicker(evolutionSpeed)
	canvasWin = tsApp.NewWindow("triangles")
	canvasWin.Resize(fyne.NewSize(side+100, side+100))
	ctrlWin = tsApp.NewWindow("buttons")
	ctrlWin.Resize(fyne.NewSize(400, 400))
	canvasCont = container.NewWithoutLayout()
}

func maxPopulationEntry() *widget.Entry {
	mp := widget.NewEntry()
	mp.SetText(strconv.Itoa(maxPopulation))
	mp.OnChanged = func(s string) {
	}

	mp.OnSubmitted = func(s string) {
		newVal, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		maxPopulation = newVal
		mp.FocusLost()
	}

	return mp
}

func startTsEntry() *widget.Entry {
	e := widget.NewEntry()
	e.SetText(strconv.Itoa(startTsCnt))
	e.OnChanged = func(s string) {
		log.Println("changed")
	}

	e.OnSubmitted = func(s string) {
		newVal, err := strconv.Atoi(s)
		if err != nil {
			startTsCnt = startTsCntDefault
			e.SetText(strconv.Itoa(startTsCnt))

		}
		startTsCnt = newVal
		log.Println("submitted")
	}
	return e
}

func main() {
	initialization()
	inputForm = widget.NewForm(widget.NewFormItem("Population", maxPopulationEntry()),
		widget.NewFormItem("start triangles count (numeric)", startTsEntry()))

	ctrlWin.SetContent(
		container.NewVBox(
			inputForm, refreshButton(),
			genRandomButton(),
			startEvolutionButton(),
			stopEvolutionButton(),
		),
	)

	canvasWin.SetContent(canvasCont)

	canvasWin.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})

	canvasWin.Show()
	ctrlWin.ShowAndRun()
}
