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

func updateTime(l *canvas.Line) {
	pos1 := fyne.Position{rand.Float32() * 400, rand.Float32() * 400}
	pos2 := fyne.Position{rand.Float32() * 400, rand.Float32() * 400}
	l.Position1 = pos2
	l.Position2 = pos1
	l.StrokeColor = color.NRGBA{0, 255, 0, 255}
	//l.Hide()
	//l.Show()
	l.Refresh()
}

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

/*func updateLineToRandom(l *canvas.Line){
	linex := canvas.NewLine(color.Black)
	l.po
}*/

//func main() {
//	a := app.New()
//	w := a.NewWindow("123")
//	w.Resize(fyne.NewSize(400, 400))
//
//	linex := canvas.NewLine(color.Black)
//	linex.StrokeColor = color.NRGBA{255, 0, 0, 255}
//	linex.StrokeWidth = 4
//	pos1 := fyne.Position{0, 0}
//	pos2 := fyne.Position{100, 100}
//	linex.Position1 = pos1
//	linex.Position2 = pos2
//
//	liney := canvas.NewLine(color.Black)
//	liney.StrokeColor = color.NRGBA{255, 0, 0, 255}
//	liney.StrokeWidth = 4
//	posy1 := fyne.Position{0, 100}
//	posy2 := fyne.Position{100, 0}
//	liney.Position1 = posy1
//	liney.Position2 = posy2
//
//	clock := widget.NewLabel("")
//	updateTime(clock)
//
//	cont := container.NewWithoutLayout(clock, linex, liney)
//
//	go func() {
//		for range time.Tick(time.Second) {
//			updateTime(clock)
//		}
//	}()
//
//	w.ShowAndRun()
//
//	w.SetContent(cont)
//
//	w.ShowAndRun()
//
//}

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
	//linex.Hide()
	rand.Seed(time.Now().UnixNano())
	//string(rand.Intn(2))
	//label := widget.NewLabel("123")

	//updateTime(linex)
	cont := container.NewWithoutLayout(linex)
	button := widget.NewButton("push me", func() {
		log.Println(123)

		updateContent(cont)
		//linex.Position2 = fyne.Position{linex.Position2.X, linex.Position2.Y + 10}
		//linex.StrokeWidth = 10
		//linex.Hide()
		//linex.Show()
	})
	//linex.Hide()

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
