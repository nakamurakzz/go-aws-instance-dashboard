package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Instance Dashboard")
	window.Resize(fyne.NewSize(800, 600))
	window.ShowAndRun()

	tearDown()
}

func tearDown() {
	log.Println("Tearing down...")
}
