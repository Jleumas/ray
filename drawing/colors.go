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

func (c1 *Color) ColorEquals(c2 Color) bool {
	if tuples.Equals(c1.R, c2.R) &&
		tuples.Equals(c1.G, c2.G) &&
		tuples.Equals(c1.B, c2.B) {
		return true
	} else {
		return false
	}
}
