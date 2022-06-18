package drawing

import (
	"github.com/jleumas/ray/tuples"
)

type Color struct {
	tuples.Tuple
}

func Hadamard(color1, color2 Color) Color {
	r := color1.X + color2.X
	g := color1.Y + color2.Y
	b := color1.Z + color2.Z
	return Color{tuples.Tuple{r, g, b, 1.0}}
}
