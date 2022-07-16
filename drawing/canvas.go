package drawing

type Position struct {
	X, Y int
}

type Canvas struct {
	height, width int
	pixels        [][]Color
}

func (canvas *Canvas) CreateCanvas(height, width int, pixels [][]Color) error {
	if pixels == nil {
		canvas.pixels = make([][]Color, height)
		for i, _ := range canvas.pixels {
			canvas.pixels[i] = make([]Color, width)
			for j := 0; j < len(canvas.pixels[i]); j++ {
				canvas.WritePixel(Color{0.0, 0.0, 0.0}, Position{i, j})
			}
		}
	} else {
		canvas.height = len(pixels)
		canvas.width = len(pixels[0])
		canvas.pixels = pixels
	}
	return nil
}

func (c *Canvas) WritePixel(col Color, pos Position) {
	c.pixels[pos.X][pos.Y] = col
}
