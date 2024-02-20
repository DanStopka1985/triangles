package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func genRandomButton() *widget.Button {
	return widget.NewButton("gen random triangles", func() {
		for i := 0; i < startTsCnt; i++ {
			addNewRandomTriangle()
		}
		for i := 0; i < startTsCnt; i++ {
			addTriangleToFyneContainer(canvasCont, aliveTs[i])
		}
		showTs(canvasCont, aliveTs)
	})
}

func genMiniTsButton() *widget.Button {
	return widget.NewButton("gen mini triangles", func() {
		for x := 0; x < 400-20; x += 15 {
			for y := 0; y < 400-20; y += 15 {
				tmp := [6]float32{10 + float32(x), 10 + float32(y), 10 + float32(x), 20 + float32(y), 20 + float32(x), 20 + float32(y)}
				addNewTriangle(tmp, _color[rand.Intn(6)])
			}
		}
		cnt := len(aliveTs)
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

func useBorder() *widget.Check {
	return widget.NewCheck("", func(b bool) {
		if b == false {
			visible_border = false

		} else {
			visible_border = true
		}
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

	inputForm = widget.NewForm(widget.NewFormItem("show border", useBorder()),
		widget.NewFormItem("Population", maxPopulationEntry()),
		widget.NewFormItem("start triangles count (numeric)", startTsEntry()))

	ctrlWin.SetContent(
		container.NewVBox(
			inputForm, refreshButton(),
			genRandomButton(),
			genMiniTsButton(),
			startEvolutionButton(),
			stopEvolutionButton(),
		),
	)
	ctrlWin.Show()
}
