package settings

import (
	"github/nakamurakzz/go-aws-instance-dashboard/src/constants"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewSettingsUi() *fyne.Container {
	appColors := constants.NewAppColors()
	log.Print(appColors.Primary, color.Black)
	sample := canvas.NewText("Settings", appColors.Primary)
	sample2 := canvas.NewText("Settings2", color.Black)
	settingsContent := container.New(
		layout.NewVBoxLayout(),
		layout.NewSpacer(),
		sample,
		layout.NewSpacer(),
		sample2,
		layout.NewSpacer(),
	)

	return settingsContent
}
