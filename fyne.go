package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

var _color = []color.NRGBA{
	{255, 0, 0, 255},
	{255, 165, 0, 255},
	{255, 255, 0, 255},
	{0, 255, 0, 255},
	{224, 255, 255, 255},
	{0, 0, 255, 255},
	{0, 155, 255, 255},
}

func showTs(c *fyne.Container, ts []triangle) {
	c.RemoveAll()
	drawBorder()

	for i := 0; i < len(ts); i++ {
		addTriangleToFyneContainer(c, ts[i])
	}

	c.Refresh()
}

func addTriangleToFyneContainer(c *fyne.Container, t triangle) {
	__color := t.color
	pos1 := fyne.Position{t.genes[0] + 20, t.genes[1] + 20}
	pos2 := fyne.Position{t.genes[2] + 20, t.genes[3] + 20}
	pos3 := fyne.Position{t.genes[4] + 20, t.genes[5] + 20}

	line1 := canvas.NewLine(color.Black)
	line1.StrokeColor = __color
	line1.StrokeWidth = 1
	line1.Position1 = pos1
	line1.Position2 = pos2

	line2 := canvas.NewLine(color.Black)
	line2.StrokeColor = __color
	line2.StrokeWidth = 1
	line2.Position1 = pos2
	line2.Position2 = pos3

	line3 := canvas.NewLine(color.Black)
	line3.StrokeColor = __color
	line3.StrokeWidth = 1
	line3.Position1 = pos3
	line3.Position2 = pos1

	c.Add(line1)
	c.Add(line2)
	c.Add(line3)
}
