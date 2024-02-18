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
	canvasCont   *fyne.Container

	ticker *time.Ticker
)

func initialization() {
	tsApp = app.New()
	rand.Seed(time.Now().UnixNano())
	ticker = time.NewTicker(evolutionSpeed)
	canvasWindow = tsApp.NewWindow("triangles")
	canvasWindow.Resize(fyne.NewSize(side+100, side+100))

}

func main() {
	initialization()

	ww := tsApp.NewWindow("buttons")

	canvasCont = container.NewWithoutLayout()

	genRandButton := widget.NewButton("gen random triangles", func() {
		ticker.Stop()
		cnt := startTsCnt
		for i := 0; i < cnt; i++ {
			addNewRandomTriangle()
		}
		for i := 0; i < cnt; i++ {
			addTriangleToFyneContainer(canvasCont, aliveTs[i])
		}
		showTs(canvasCont, aliveTs)
	})

	/*	add1RandomButton := widget.NewButton("add 1 random", func() {
		addNewRandomTriangle()
		showTs(canvasCont, aliveTs)
	})*/

	//killLastTriangleButton := widget.NewButton("kill last", func() {
	//	sortAliveTs()
	//	killLastTriangle()
	//	showTs(canvasCont, aliveTs)
	//})

	/*	showDeathTsButton := widget.NewButton("show death", func() {
			showTs(canvasCont, deathTs)
		})

		showAliveTsButton := widget.NewButton("show alive", func() {
			showTs(canvasCont, aliveTs)
		})*/

	/*	naturalSelectionTsButton := widget.NewButton("natural selection (100)", func() {
			naturalSelection()
			showTs(canvasCont, aliveTs)
		})

		newGenerationButton := widget.NewButton("new generation", func() {
			createNewGeneration()
			showTs(canvasCont, aliveTs)
		})*/

	/*	loop20 := widget.NewButton("loop 20 generation selection", func() {
		for i := 0; i < 2000000; i++ {
			createNewGeneration()
			naturalSelection()
			time.Sleep(50 * time.Millisecond)
			showTs(canvasCont, aliveTs)
		}
		//showTs(canvasCont, aliveTs)
		max := float64(0)
		ii := -1
		for i := 0; i < len(aliveTs); i++ {
			if max < aliveTs[i].power {
				max = aliveTs[i].power
				ii = 1
			}

		}
		log.Println(max)
		log.Println(aliveTs[ii])
		//showTs(canvasCont, []triangle{aliveTs[ii]})
	})*/

	refresh := widget.NewButton("refresh", func() {
		aliveTs = nil

		showTs(canvasCont, aliveTs)
	})

	startEvolution := widget.NewButton("start evolution", func() {
		ticker = time.NewTicker(evolutionSpeed)
		go func() {
			for range ticker.C {

				createNewGeneration()
				naturalSelection()
				showTs(canvasCont, aliveTs)
			}
		}()
	})

	stopEvolution := widget.NewButton("stop evolution", func() {
		ticker.Stop()
	})

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

	ww.SetContent(container.NewVBox(form, refresh, genRandButton, startEvolution, stopEvolution))

	canvasWindow.SetContent(canvasCont)

	//go func() {
	//	for range time.Tick(time.Second) {
	//		updateTime(clock)
	//	}
	//}()
	canvasWindow.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})

	ww.Resize(fyne.NewSize(400, 400))
	canvasWindow.Show()
	ww.ShowAndRun()
}
