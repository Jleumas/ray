package drawing

import (
	"testing"
)

func TestHadamard(t *testing.T) {
	var testColors = []struct {
		A, B Color
	}{
		{Color{1.0, 0.2, 0.4}, Color{0.9, 0.2, 0.04}},
		{Color{0.23, 0.34, 0.45}, Color{0.56, 0.67, 0.78}},
		{Color{}, Color{}},
		{Color{}, Color{}},
		{Color{}, Color{}},
	}

	for _, testC := range testColors {

	}
}
