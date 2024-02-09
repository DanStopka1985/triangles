package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"math/rand"
	"time"
)

var (
	aliveTs []triangle
	deathTs []triangle
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("triangles")
	w.Resize(fyne.NewSize(400, 400))

	ww := a.NewWindow("buttons")

	cont := container.NewWithoutLayout()
	gen5randButton := widget.NewButton("gen 5 random triangles", func() {
		cnt := 102
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

	ww.SetContent(container.NewVBox(gen5randButton, add1RandomButton, killLastTriangleButton, showDeathTsButton, showAliveTsButton, naturalSelectionTsButton))

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
