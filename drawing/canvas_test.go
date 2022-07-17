package drawing

import (
	"testing"
)

func TestCanvasEquals(t *testing.T) {

}

func TestCreateCanvas(t *testing.T) {
	black := Color{1.0, 1.0, 1.0}
	red := Color{1.0, 0.0, 0.0}
	var testCanvas = []struct {
		len, wid int
		result   Canvas
	}{
		{4, 5, Canvas{4, 4, nil}},
		{3, 2, Canvas{3, 2, [][]Color{
			{red, black, red},
			{black, red, black},
			{red, black, red},
		}}},
	}

	for _, testC := range testCanvas {
		t.Run("Compute Hadamard product of Colors", func(t *testing.T) {
			canvas := Canvas{}
			err := canvas.CreateCanvas(testC.len, testC.wid, [][]Color{
				{red, black, red},
				{black, red, black},
				{red, black, red},
			})
			if err != nil && !CanvasEquals(canvas, testC.result) {
				t.Errorf("Not working! Canvases are not equal")
			}
		})
	}
}
