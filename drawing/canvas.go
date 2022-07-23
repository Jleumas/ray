package drawing

type Position struct {
	X, Y int
}

type Canvas struct {
	height, width int
	pixels        [][]Color
}

func (canvas *Canvas) Init(height, width int) (error, Canvas) {
	canvas.pixels = make([][]Color, height)
	for i := 0; i < len(canvas.pixels[0]); i++ {
		canvas.pixels[i] = make([]Color, width)
	}

	return nil, *canvas
}

func (canvas *Canvas) CanvasEquals(other *Canvas) bool {
	if canvas.height != other.height {
		return false
	}
	if canvas.width != other.width {
		return false
	}
	if canvas.pixels == nil || other.pixels == nil {
		return false
	}
	for i := 0; i < len(canvas.pixels); i++ {
		for j := 0; j < len(canvas.pixels[0]); j++ {
			if !ColorEquals(canvas.pixels[i][j], other.pixels[i][j]) {
				return false
			}
		}
	}

	return true
}

func (canvas *Canvas) WritePixel(col Color, pos Position) {
	canvas.pixels[pos.X][pos.Y] = col
}

func (canvas *Canvas) GetPixel(p Position) Color {
	return canvas.pixels[p.X][p.Y]
}
