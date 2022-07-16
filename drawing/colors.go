package drawing

import "github.com/jleumas/ray/tuples"

type Color struct {
	R, G, B float64
}

func Hadamard(colors ...Color) Color {
	color := Color{1.0, 1.0, 1.0}

	for _, c := range colors {
		color.R *= c.R
		color.G *= c.G
		color.B *= c.B
	}

	return color
}

func ColorEquals(c1, c2 Color) bool {
	if tuples.Equals(c1.R, c2.R) &&
		tuples.Equals(c1.G, c2.G) &&
		tuples.Equals(c1.B, c2.B) {
		return true
	} else {
		return false
	}
}

func Add(colors ...Color) Color {
	newC := Color{}

	for _, c := range colors {
		newC.R += c.R
		newC.G += c.G
		newC.B += c.B
	}
	return newC
}

func Subtract(c1, c2 Color) Color {
	newC := Color{}
	newC.R = c1.R - c2.R
	newC.G = c1.G - c2.G
	newC.B = c1.B - c2.B

	return newC
}

func MultiplyScalar(c Color, scalar float64) Color {
	newC := Color{}
	newC.R = c.R * scalar
	newC.G = c.G * scalar
	newC.B = c.B * scalar

	return newC
}
