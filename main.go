package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"math/rand"
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
	ctrlWinInit()
	canvasWinInit()
}

func main() {
	initialization()
	tsApp.Run()
}
