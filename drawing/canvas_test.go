package drawing

import (
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	black := Color{1.0, 1.0, 1.0}
	red:= Color{1.0, 0.0, 0.0}
	var testCanvas = []struct {
		len, wid int
		result   Canvas
	}{
		{4, 5, Canvas{4, 4, nil}},
		{3, 2, Canvas{3, 2, [][]Color{
			{black, black, black},
			{black, black, black},
			{black, black, black}
		}}},
	}
}
