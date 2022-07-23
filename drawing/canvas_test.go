package drawing

import (
	"testing"
)

func TestWritePixel(t *testing.T) {
	var testCanvas = []struct {
		canvas          Canvas
		position        Position
		result          Color
		testDescription string
	}{
		{Canvas{1, 1, make([][]Color, 1)}, Position{0, 0}, Color{1.0, 1.0, 1.0}, "Tests white pixel is white"},
		{Canvas{1, 1, make([][]Color, 1)}, Position{0, 0}, Color{0.0, 0.0, 0.0}, "Tests black pixel is black"},
	}

	for _, testC := range testCanvas {
		t.Run(testC.testDescription, func(t *testing.T) {
			testC.canvas.pixels[0] = make([]Color, 1)
			testC.canvas.WritePixel(testC.result, testC.position)
			if !(testC.canvas.GetPixel(testC.position) == testC.result) {
				t.Errorf("Not working! Failure: %v", testC.testDescription)
			}
		})
	}
}

func TestCanvasEquals(t *testing.T) {
	black := Color{1.0, 1.0, 1.0}
	red := Color{1.0, 0.0, 0.0}
	var testCanvas = []struct {
		A, B            Canvas
		result          bool
		testDescription string
	}{
		{Canvas{4, 4, nil}, Canvas{4, 4, nil}, false, "Two canvases with uninitalized pixel fields cannot be equal"},
		{Canvas{3, 2, [][]Color{
			{red, black, red},
			{black, red, black},
			{red, black, red},
		}}, Canvas{3, 2, nil}, false, "Test that different canvases are not equal"},
		{Canvas{3, 2, [][]Color{
			{red, black, red},
			{black, red, black},
			{red, black, red},
		}}, Canvas{3, 2, [][]Color{
			{red, black, red},
			{black, red, black},
			{red, black, red},
		}}, true, "Identical canvases are equal"},
	}

	for _, testC := range testCanvas {
		t.Run(testC.testDescription, func(t *testing.T) {
			if !(CanvasEquals(&testC.A, &testC.B) == testC.result) {
				t.Errorf("Not working! Failure: %v", testC.testDescription)
			}
		})
	}
}
