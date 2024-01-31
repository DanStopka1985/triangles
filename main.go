package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"log"
	"math/rand"
	"time"
)

type myEntry struct {
	widget.Entry
}

func updateTime(clock *widget.Label) int {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
	return 0
}

func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}

	log.Println("Shortcut typed:", s)
}

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
	ww := a.NewWindow("button")

	rand.Seed(time.Now().UnixNano())
	//string(rand.Intn(2))
	//label := widget.NewLabel("123")
	clock := widget.NewLabel("")
	button := widget.NewButton("push me", func() {
		log.Println(123)

		updateTime(clock)
	})
	cont := container.NewWithoutLayout(clock)
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
