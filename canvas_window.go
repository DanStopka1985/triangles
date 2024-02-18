package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"log"
	"os"
)

func canvasWinInit() {
	canvasWin = tsApp.NewWindow("triangles")
	canvasWin.Resize(fyne.NewSize(side+100, side+100))
	canvasCont = container.NewWithoutLayout()
	canvasWin.SetContent(canvasCont)

	canvasWin.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})
	canvasWin.Show()

}
