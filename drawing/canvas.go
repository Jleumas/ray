package drawing

type Position struct {
	X, Y int
}

type Canvas struct {
	height, width int
	pixels        [][]Color
}

func CanvasEquals(a, b *Canvas) bool {
	if a.height != b.height {
		return false
	}
	if a.width != b.width {
		return false
	}
	if a.pixels == nil || b.pixels == nil {
		return false
	}
	for i := 0; i < len(a.pixels); i++ {
		for j := 0; j < len(a.pixels[0]); j++ {
			if !ColorEquals(a.pixels[i][j], b.pixels[i][j]) {
				return false
			}
		}
	}

	return true
}

func (c *Canvas) WritePixel(col Color, pos Position) {
	c.pixels[pos.X][pos.Y] = col
}
