package main

import (
	"fyne.io/fyne/v2/widget"
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
