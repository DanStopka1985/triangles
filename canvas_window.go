package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"log"
	"os"
)

func drawBorder() {
	if visible_border == false {
		return
	}
	r := canvas.NewRectangle(color.White)
	r.StrokeColor = color.Black
	r.StrokeWidth = 1
	r.Resize(fyne.NewSize(side+4, side+4))
	r.Move(fyne.Position{X: 19 - 2, Y: 19 - 2})
	canvasCont.Add(r)
}

func canvasWinInit() {

	canvasWin = tsApp.NewWindow("triangles")
	canvasWin.Resize(fyne.NewSize(side+100, side+100))
	canvasCont = container.NewWithoutLayout()
	drawBorder()
	canvasWin.SetContent(canvasCont)

	canvasWin.SetOnClosed(
		func() {
			log.Println("window closed")
			os.Exit(0)
		})
	canvasWin.Show()

}
