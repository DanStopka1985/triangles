package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("triangles")
	w.Resize(fyne.NewSize(400, 400))

	ww := a.NewWindow("buttons")

	cont := container.NewWithoutLayout()
	gen5randButton := widget.NewButton("gen 5 random triangles", func() {
		cnt := 5
		for i := 0; i < cnt; i++ {
			addNewRandomTriangle()
		}
		for i := 0; i < cnt; i++ {
			addTriangleToFyneContainer(cont, aliveTs[i])
		}
		showTs(cont, aliveTs)
	})

	add1RandomButton := widget.NewButton("add 1 random", func() {
		addNewRandomTriangle()
		showTs(cont, aliveTs)
	})

	killLastTriangleButton := widget.NewButton("kill last", func() {
		sortAliveTs()
		killLastTriangle()
		showTs(cont, aliveTs)
	})

	showDeathTsButton := widget.NewButton("show death", func() {
		showTs(cont, deathTs)
	})

	showAliveTsButton := widget.NewButton("show alive", func() {
		showTs(cont, aliveTs)
	})

	naturalSelectionTsButton := widget.NewButton("natural selection (100)", func() {
		naturalSelection()
		showTs(cont, aliveTs)
	})

	newGenerationButton := widget.NewButton("new generation", func() {
		createNewGeneration()
		showTs(cont, aliveTs)
	})

	loop20 := widget.NewButton("loop 20 generation selection", func() {
		for i := 0; i < 20; i++ {
			createNewGeneration()
			naturalSelection()
		}
		showTs(cont, aliveTs)
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
	})

	refresh := widget.NewButton("refresh", func() {
		aliveTs = nil

		showTs(cont, aliveTs)
	})

	ww.SetContent(container.NewVBox(refresh, gen5randButton, add1RandomButton, killLastTriangleButton, showDeathTsButton,
		showAliveTsButton, naturalSelectionTsButton, newGenerationButton, loop20))

	w.SetContent(cont)

	//go func() {
	//	for range time.Tick(time.Second) {
	//		updateTime(clock)
	//	}
	//}()

	ww.Resize(fyne.NewSize(400, 400))
	ww.Show()
	w.ShowAndRun()
}
