package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var (
	tsApp      fyne.App
	canvasWin  fyne.Window
	ctrlWin    fyne.Window
	canvasCont *fyne.Container
	inputForm  *widget.Form
)

func uiInit() {
	ctrlWinInit()
	canvasWinInit()
}
