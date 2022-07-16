package drawing

type Position struct {
	X, Y int
}

type Canvas struct {
	Position
	pixels [][]Color
	p      **Color
}

func (c *Canvas) WritePixel(col Color, pos Position) {
	c.pixels[pos.X][pos.Y] = col
}
