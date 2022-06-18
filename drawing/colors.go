package drawing

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
