package constants

import "image/color"

type AppColors struct {
	Background     color.Color
	Primary        color.Color
	PrimaryBariant color.Color
	Secondary      color.Color
	Font           color.Color
	Success        color.Color
	Error          color.Color
}

func NewAppColors() *AppColors {
	appColors := AppColors{
		// Background: #fffffe
		Background: color.RGBA{255, 255, 254, 255},
		// Primary: #ff595e
		Primary: color.RGBA{255, 89, 94, 255},
		// PrimaryBariant: #ffca3a
		PrimaryBariant: color.RGBA{255, 202, 58, 255},
		// Secondary: #8ac926
		Secondary: color.RGBA{138, 201, 38, 255},
		// Font: #1d3557
		Font: color.RGBA{29, 53, 87, 255},
		// Success: #1982c4
		Success: color.RGBA{25, 130, 196, 255},
		// Error: #6a4c93
		Error: color.RGBA{106, 76, 147, 255},
	}
	return &appColors
}
