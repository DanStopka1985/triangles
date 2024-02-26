package main

import (
	"fyne.io/fyne/v2/app"
	"math/rand"
	"time"
)

var (
	ticker *time.Ticker
)

func initialization() {
	tsApp = app.New()
	rand.Seed(time.Now().UnixNano())
	ticker = time.NewTicker(evolutionSpeed)
	uiInit()
}

func main() {

	initialization()

	getChanceTest()
	tsApp.Run()
}
