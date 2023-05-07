package constants

import (
	"image/color"
	"testing"
)

func TestNewAppColors(t *testing.T) {
	appColors := NewAppColors()

	if appColors.Background == nil {
		t.Errorf("Expected appColors.Background to be non-nil")
	}

	if appColors.Primary == nil {
		t.Errorf("Expected appColors.Primary to be non-nil")
	}

	if appColors.PrimaryBariant == nil {
		t.Errorf("Expected appColors.PrimaryBariant to be non-nil")
	}

	if appColors.Secondary == nil {
		t.Errorf("Expected appColors.Secondary to be non-nil")
	}

	if appColors.Font == nil {
		t.Errorf("Expected appColors.Font to be non-nil")
	}

	if appColors.Success == nil {
		t.Errorf("Expected appColors.Success to be non-nil")
	}

	if appColors.Error == nil {
		t.Errorf("Expected appColors.Error to be non-nil")
	}

	bg := color.RGBA{255, 255, 254, 255}
	if appColors.Background != bg {
		t.Errorf("Expected appColors.Background to be %v, got %v", bg, appColors.Background)
	}

	p := color.RGBA{255, 89, 94, 255}
	if appColors.Primary != p {
		t.Errorf("Expected appColors.Primary to be %v, got %v", p, appColors.Primary)
	}

	pv := color.RGBA{255, 202, 58, 255}
	if appColors.PrimaryBariant != pv {
		t.Errorf("Expected appColors.PrimaryBariant to be %v, got %v", pv, appColors.PrimaryBariant)
	}

	s := color.RGBA{138, 201, 38, 255}
	if appColors.Secondary != s {
		t.Errorf("Expected appColors.Secondary to be %v, got %v", s, appColors.Secondary)
	}

	f := color.RGBA{29, 53, 87, 255}
	if appColors.Font != f {
		t.Errorf("Expected appColors.Font to be %v, got %v", f, appColors.Font)
	}

	su := color.RGBA{25, 130, 196, 255}
	if appColors.Success != su {
		t.Errorf("Expected appColors.Success to be %v, got %v", su, appColors.Success)
	}

	err := color.RGBA{106, 76, 147, 255}
	if appColors.Error != err {
		t.Errorf("Expected appColors.Error to be %v, got %v", err, appColors.Error)
	}
}
