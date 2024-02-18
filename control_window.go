package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
	"time"
)

func genRandomButton() *widget.Button {
	return widget.NewButton("gen random triangles", func() {
		cnt := startTsCnt
		for i := 0; i < cnt; i++ {
			addNewRandomTriangle()
		}
		for i := 0; i < cnt; i++ {
			addTriangleToFyneContainer(canvasCont, aliveTs[i])
		}
		showTs(canvasCont, aliveTs)
	})
}

func refreshButton() *widget.Button {
	return widget.NewButton("refresh", func() {
		aliveTs = nil
		showTs(canvasCont, aliveTs)
	})
}

func startEvolutionButton() *widget.Button {
	return widget.NewButton("start evolution", func() {
		ticker = time.NewTicker(evolutionSpeed)
		go func() {
			for range ticker.C {
				createNewGeneration()
				naturalSelection()
				showTs(canvasCont, aliveTs)
			}
		}()
	})
}

func stopEvolutionButton() *widget.Button {
	return widget.NewButton("stop evolution", func() {
		ticker.Stop()
	})
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

func ctrlWinInit() {
	ctrlWin = tsApp.NewWindow("buttons")
	ctrlWin.Resize(fyne.NewSize(400, 400))

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
	ctrlWin.Show()
}
