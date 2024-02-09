package main

import "fyne.io/fyne/v2"

func showTs(c *fyne.Container, ts []triangle) {
	c.RemoveAll()

	for i := 0; i < len(ts); i++ {
		addTriangleToFyneContainer(c, ts[i])
	}

	c.Refresh()
}
