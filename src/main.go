package main

import (
	"github/nakamurakzz/go-aws-instance-dashboard/src/settings"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Instance Dashboard")
	window.Resize(fyne.NewSize(800, 600))

	settingsContent := settings.NewSettingsUi()
	window.SetContent(settingsContent)

	window.ShowAndRun()

	tearDown()
}

func tearDown() {
	log.Println("Tearing down...")
}
