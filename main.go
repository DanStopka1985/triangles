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

func main() {
	rand.Seed(time.Now().UnixNano())
	Ticker := time.NewTicker(evolutionSpeed)
	a := app.New()
	w := a.NewWindow("triangles")
	w.Resize(fyne.NewSize(side+100, side+100))
	//w.SetFullScreen(true)

	ww := a.NewWindow("buttons")

	cont := container.NewWithoutLayout()
	genRandButton := widget.NewButton("gen random triangles", func() {
		//Ticker.Stop()
		cnt := startTsCnt
		for i := 0; i < cnt; i++ {
			addNewRandomTriangle()
		}
		for i := 0; i < cnt; i++ {
			addTriangleToFyneContainer(cont, aliveTs[i])
		}
		showTs(cont, aliveTs)
	})

	/*	add1RandomButton := widget.NewButton("add 1 random", func() {
		addNewRandomTriangle()
		showTs(cont, aliveTs)
	})*/

	//killLastTriangleButton := widget.NewButton("kill last", func() {
	//	sortAliveTs()
	//	killLastTriangle()
	//	showTs(cont, aliveTs)
	//})

	/*	showDeathTsButton := widget.NewButton("show death", func() {
			showTs(cont, deathTs)
		})

		showAliveTsButton := widget.NewButton("show alive", func() {
			showTs(cont, aliveTs)
		})*/

	/*	naturalSelectionTsButton := widget.NewButton("natural selection (100)", func() {
			naturalSelection()
			showTs(cont, aliveTs)
		})

		newGenerationButton := widget.NewButton("new generation", func() {
			createNewGeneration()
			showTs(cont, aliveTs)
		})*/

	/*	loop20 := widget.NewButton("loop 20 generation selection", func() {
		for i := 0; i < 2000000; i++ {
			createNewGeneration()
			naturalSelection()
			time.Sleep(50 * time.Millisecond)
			showTs(cont, aliveTs)
		}
		//showTs(cont, aliveTs)
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
		//showTs(cont, []triangle{aliveTs[ii]})
	})*/

	refresh := widget.NewButton("refresh", func() {
		aliveTs = nil

		showTs(cont, aliveTs)
	})

	startEvolution := widget.NewButton("start evolution", func() {
		Ticker = time.NewTicker(evolutionSpeed)
		go func() {
			for range Ticker.C {

				createNewGeneration()
				naturalSelection()
				showTs(cont, aliveTs)
			}
		}()
	})

	stopEvolution := widget.NewButton("stop evolution", func() {
		Ticker.Stop()
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

	ww.SetContent(container.NewVBox(form, refresh, genRandButton, /*, add1RandomButton, showDeathTsButton,
		showAliveTsButton, naturalSelectionTsButton, newGenerationButton, loop20*/startEvolution, stopEvolution))

	w.SetContent(cont)

	//go func() {
	//	for range time.Tick(time.Second) {
	//		updateTime(clock)
	//	}
	//}()
	w.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})

	ww.Resize(fyne.NewSize(400, 400))
	w.Show()
	ww.ShowAndRun()
}
