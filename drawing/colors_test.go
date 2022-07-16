package drawing

import (
	"testing"
)

func TestHadamard(t *testing.T) {
	var testColors = []struct {
		A, B, result Color
	}{
		{Color{1.0, 0.2, 0.4}, Color{0.9, 1, 0.1}, Color{0.9, 0.2, 0.04}},
	}

	for _, testC := range testColors {
		t.Run("Hadamard", func(t *testing.T) {
			product := Hadamard(testC.A, testC.B)
			if !ColorEquals(product, testC.result) {
				t.Errorf("Not working! Color{ %v, %v, %v} is product",
					product.R, product.G, product.B)
			}
		})
	}
}

func TestAddColors(t *testing.T) {
	var testColors = []struct {
		A, B, result Color
	}{
		{Color{0.9, 0.6, 0.75}, Color{0.7, 0.1, 0.25}, Color{1.6, 0.7, 1.0}},
	}
	for _, testC := range testColors {
		t.Run("Add", func(t *testing.T) {
			product := Add(testC.A, testC.B)
			if !ColorEquals(product, testC.result) {
				t.Errorf("Not working. Color{ %v, %v, %v} is product",
					product.R, product.G, product.B)
			}
		})
	}
}

func TestSubtractColors(t *testing.T) {
	var testColors = []struct {
		A, B, result Color
	}{
		{Color{0.9, 0.6, 0.75}, Color{0.7, 0.1, 0.25}, Color{0.2, 0.5, 0.5}},
	}
	for _, testC := range testColors {
		t.Run("Subtract", func(t *testing.T) {
			product := Subtract(testC.A, testC.B)
			if !ColorEquals(product, testC.result) {
				t.Errorf("Not working. Color{ %v, %v, %v} is product",
					product.R, product.G, product.B)
			}
		})
	}
}

func TestMultiplyColors(t *testing.T) {
	var testColors = []struct {
		A      Color
		scalar float64
		result Color
	}{
		{Color{0.2, 0.3, 0.4}, 2.0, Color{0.4, 0.6, 0.8}},
	}
	for _, testC := range testColors {
		t.Run("Multiply", func(t *testing.T) {
			product := Multiply(testC.A, testC.scalar)
			if !ColorEquals(product, testC.result) {
				t.Errorf("Not working. Color{ %v, %v, %v} is product",
					product.R, product.G, product.B)
			}
		})
	}
}
