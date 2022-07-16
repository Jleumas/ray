package drawing

import (
	"testing"
)

func TestHadamard(t *testing.T) {
	var testColors = []struct {
		A, B, answer Color
	}{
		{Color{1.0, 0.2, 0.4}, Color{0.9, 1, 0.1}, Color{0.9, 0.2, 0.04}},
	}

	for _, testC := range testColors {
		t.Run("Hadamard", func(t *testing.T) {
			product := Hadamard(testC.A, testC.B)
			if !(product.ColorEquals(testC.answer)) {
				t.Errorf("Not working! Color{ %v, %v, %v} is product",
					product.R, product.G, product.B)
			}
		})
	}
}
