package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"
	"time"
)

var (
	aliveTs []triangle
	deathTs []triangle
)

func addTriangleToFyneContainer(c *fyne.Container, t triangle) {
	pos1 := fyne.Position{t.coordinates[0], t.coordinates[1]}
	pos2 := fyne.Position{t.coordinates[2], t.coordinates[3]}
	pos3 := fyne.Position{t.coordinates[4], t.coordinates[5]}

	line1 := canvas.NewLine(color.Black)
	line1.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line1.StrokeWidth = 1
	line1.Position1 = pos1
	line1.Position2 = pos2

	line2 := canvas.NewLine(color.Black)
	line2.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line2.StrokeWidth = 1
	line2.Position1 = pos2
	line2.Position2 = pos3

	line3 := canvas.NewLine(color.Black)
	line3.StrokeColor = color.NRGBA{255, 0, 0, 255}
	line3.StrokeWidth = 1
	line3.Position1 = pos3
	line3.Position2 = pos1

	c.Add(line1)
	c.Add(line2)
	c.Add(line3)
}

func addNewTrianglesAndShow(c *fyne.Container) {
	c.RemoveAll()
	cnt := 3
	aliveTs = genRandomTriangles(cnt)

	for i := 0; i < cnt; i++ {
		addTriangleToFyneContainer(c, aliveTs[i])
	}

	c.Refresh()
}

func addNew1TriangleAndShow(c *fyne.Container) {
	c.RemoveAll()
	aliveTs = append(aliveTs, genRandomTriangle())

	for i := 0; i < len(aliveTs); i++ {
		addTriangleToFyneContainer(c, aliveTs[i])
	}

	c.Refresh()
}

func killLastTriangle() {
	if len(aliveTs) == 0 {
		return
	}

	deathTs = append(deathTs, aliveTs[len(aliveTs)-1])

	if len(aliveTs) > 0 {
		aliveTs = aliveTs[:len(aliveTs)-1]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	a := app.New()
	w := a.NewWindow("triangles")
	w.Resize(fyne.NewSize(400, 400))

	ww := a.NewWindow("buttons")

	cont := container.NewWithoutLayout()
	gen5randButton := widget.NewButton("gen 5 random triangles", func() {
		addNewTrianglesAndShow(cont)
	})

	add1RandomButton := widget.NewButton("add 1 random", func() {
		addNew1TriangleAndShow(cont)
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

	ww.SetContent(container.NewVBox(gen5randButton, add1RandomButton, killLastTriangleButton, showDeathTsButton, showAliveTsButton))

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
