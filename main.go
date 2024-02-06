package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"math/rand"
	"time"
)

func updateContent(c *fyne.Container) {
	c.RemoveAll()
	pos1 := fyne.Position{rand.Float32() * 400, rand.Float32() * 400}
	pos2 := fyne.Position{rand.Float32() * 400, rand.Float32() * 400}
	pos3 := fyne.Position{rand.Float32() * 400, rand.Float32() * 400}

	line1 := canvas.NewLine(color.Black)
	line1.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line1.StrokeWidth = 4
	line1.Position1 = pos1
	line1.Position2 = pos2

	line2 := canvas.NewLine(color.Black)
	line2.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line2.StrokeWidth = 4
	line2.Position1 = pos2
	line2.Position2 = pos3

	line3 := canvas.NewLine(color.Black)
	line3.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line3.StrokeWidth = 4
	line3.Position1 = pos3
	line3.Position2 = pos1

	c.Add(line1)
	c.Add(line2)
	c.Add(line3)
	c.Refresh()
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	w.Resize(fyne.NewSize(400, 400))

	linex := canvas.NewLine(color.Black)
	linex.StrokeColor = color.NRGBA{255, 0, 0, 255}
	linex.StrokeWidth = 4
	pos1 := fyne.Position{0, 0}
	pos2 := fyne.Position{100, 100}
	linex.Position1 = pos2
	linex.Position2 = pos1

	ww := a.NewWindow("button")
	rand.Seed(time.Now().UnixNano())

	cont := container.NewWithoutLayout(linex)
	button := widget.NewButton("push me", func() {
		log.Println("pushed")

		updateContent(cont)
	})

	ww.SetContent(button)
	w.SetContent(cont)

	//go func() {
	//	for range time.Tick(time.Second) {
	//		updateTime(clock)
	//	}
	//}()

	ww.Show()
	w.ShowAndRun()
}
