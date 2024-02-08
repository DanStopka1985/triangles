package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
	"image/color"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
)

type triangle struct {
	generation  int
	coordinates [6]float32
	power       float32
	color       int
	uuid        uuid.UUID
}

var (
	aliveTs []triangle
	deathTs []triangle
)

func abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func genRandomTriangle() triangle {
	var t triangle
	t.uuid = uuid.New()
	for i := 0; i < 6; i++ {
		t.coordinates[i] = rand.Float32() * 400
	}
	temp1 := (t.coordinates[2] - t.coordinates[0]) * (t.coordinates[5] - t.coordinates[1])
	temp2 := (t.coordinates[4] - t.coordinates[0]) * (t.coordinates[3] - t.coordinates[1])

	t.power = float32(0.5 * abs(temp1-temp2))

	return t
}

func genRandomTriangles(cnt int) []triangle {
	r := make([]triangle, 0)
	for i := 0; i < cnt; i++ {
		t := genRandomTriangle()
		r = append(r, t)
	}
	return r
}

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

func killLastTriangle(c *fyne.Container) {
	if len(aliveTs) == 0 {
		return
	}
	c.RemoveAll()
	deathTs = append(deathTs, aliveTs[len(aliveTs)-1])
	log.Println(len(deathTs))
	if len(aliveTs) > 0 {
		aliveTs = aliveTs[:len(aliveTs)-1]
	}

	for i := 0; i < len(aliveTs); i++ {
		addTriangleToFyneContainer(c, aliveTs[i])
	}

	c.Refresh()
}

func showDeathTs(c *fyne.Container) {
	c.RemoveAll()
	log.Println(len(deathTs))
	for i := 0; i < len(deathTs); i++ {
		addTriangleToFyneContainer(c, deathTs[i])
	}

	c.Refresh()
}

func showAliveTs(c *fyne.Container) {
	c.RemoveAll()

	for i := 0; i < len(aliveTs); i++ {
		addTriangleToFyneContainer(c, aliveTs[i])
	}

	c.Refresh()
}

func sortAliveTs() {
	sort.Slice(aliveTs, func(i, j int) bool {
		return aliveTs[i].power > aliveTs[j].power
	})
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
		killLastTriangle(cont)
	})

	showDeathTsButton := widget.NewButton("show death", func() {
		showDeathTs(cont)
	})

	showAliveTsButton := widget.NewButton("show alive", func() {
		showAliveTs(cont)
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
