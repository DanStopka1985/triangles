package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

func showTs(c *fyne.Container, ts []triangle) {
	c.RemoveAll()

	for i := 0; i < len(ts); i++ {
		addTriangleToFyneContainer(c, ts[i])
	}

	c.Refresh()
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
