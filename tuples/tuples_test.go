package tuples

import (
	"fmt"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	var testPoints = []struct {
		X, Y, Z float64
	}{
		{0.0, 0.0, 0.0},
		{4.3, -4.2, 3.1},
		{-4.3, 4.2, -3.1},
	}

	for _, testPt := range testPoints {
		testName := fmt.Sprintf("%v, %v, %v", testPt.X, testPt.Y, testPt.Z)
		t.Run(testName, func(t *testing.T) {
			point := Point(testPt.X, testPt.Y, testPt.Z)
			if point.W != 1.0 {
				t.Errorf("Wrong w")
			}
			if point.X != testPt.X {
				t.Errorf("Wrong x")
			}
			if point.Y != testPt.Y {
				t.Errorf("Wrong y")
			}
			if point.Z != testPt.Z {
				t.Errorf("Wrong z")
			}
		})
	}
}

func TestVector(t *testing.T) {
	var testVectors = []struct {
		X, Y, Z float64
	}{
		{0.0, 0.0, 0.0},
		{4.3, -4.2, 3.1},
		{-4.3, 4.2, -3.1},
	}

	for _, testVec := range testVectors {
		testName := fmt.Sprintf("%v, %v, %v", testVec.X, testVec.Y, testVec.Z)
		t.Run(testName, func(t *testing.T) {
			point := Vector(testVec.X, testVec.Y, testVec.Z)
			if point.W != 0.0 {
				t.Errorf("Wrong w")
			}
			if point.X != testVec.X {
				t.Errorf("Wrong x")
			}
			if point.Y != testVec.Y {
				t.Errorf("Wrong y")
			}
			if point.Z != testVec.Z {
				t.Errorf("Wrong z")
			}
		})
	}
}

func TestTupleEquals(t *testing.T) {
	var testTuples = []struct {
		a, b Tuple
		ok   bool
	}{
		{Tuple{3.0, -2.0, 5.0, 1.0}, Tuple{3.0, -2.0, 5.0, 1.0}, true},
		{Tuple{-1.0, 2.0, -10.0, 0.0}, Tuple{-1.0, 2.0, -10.0, 0.0}, true},
		{Tuple{1.0, 0.0, 8.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}, false},
		{Tuple{0.0, 50.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}, false},
		{Point(1.0, 1.0, 1.0), Vector(1.0, 1.0, 1.0), false},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Comparing Tuple {%v, %v, %v, %v} to Tuple {%v, %v, %v, %v} for equality (Expected Result: %v)",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.b.X, testTup.b.Y, testTup.b.Z, testTup.b.W,
			testTup.ok)
		t.Run(testName, func(t *testing.T) {
			if TupleEquals(testTup.a, testTup.b) != testTup.ok {
				t.Errorf("Cannot correctly compare tuples using equality function!")
			}
		})
	}
}

func TestAdd(t *testing.T) {
	var testTuples = []struct {
		a, b, result Tuple
	}{
		{Tuple{3.0, -2.0, 5.0, 1.0}, Tuple{-2.0, 3.0, 1.0, 0.0}, Tuple{1.0, 1.0, 6.0, 1.0}},
		{Tuple{0.0, 0.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}},
		{Tuple{0.0, 0.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Adding Tuple {%v, %v, %v, %v} to Tuple {%v, %v, %v, %v} (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.b.X, testTup.b.Y, testTup.b.Z, testTup.b.W,
			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			if !TupleEquals(Add(testTup.a, testTup.b), testTup.result) {
				t.Errorf("Cannot correctly add tuples using standalone function!")
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	var testTuples = []struct {
		a, b, result Tuple
	}{
		{Tuple{3.0, 2.0, 1.0, 1.0}, Tuple{5.0, 6.0, 7.0, 1.0}, Tuple{-2.0, -4.0, -6.0, 0.0}},
		{Tuple{3.0, 2.0, 1.0, 1.0}, Tuple{5.0, 6.0, 7.0, 0.0}, Tuple{-2.0, -4.0, -6.0, 1.0}},
		{Tuple{3.0, 2.0, 1.0, 0.0}, Tuple{5.0, 6.0, 7.0, 0.0}, Tuple{-2.0, -4.0, -6.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Subtracting Tuple {%v, %v, %v, %v} from Tuple {%v, %v, %v, %v} (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.b.X, testTup.b.Y, testTup.b.Z, testTup.b.W,
			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			if !TupleEquals(Subtract(testTup.a, testTup.b), testTup.result) {
				t.Errorf("Cannot correctly subtract tuples using standalone function!")
			}
		})
	}
}

func TestNegate(t *testing.T) {
	var testTuples = []struct {
		a, result Tuple
	}{
		{Tuple{3.0, 2.0, 1.0, 1.0}, Tuple{-3.0, -2.0, -1.0, -1.0}},
		{Tuple{1.0, 2.0, 3.0, 4.0}, Tuple{-1.0, -2.0, -3.0, -4.0}},
		{Tuple{-1.0, -2.0, -3.0, 0.0}, Tuple{1.0, 2.0, 3.0, 0.0}},
		{Tuple{0.0, 0.0, 0.0, 0.0}, Tuple{0.0, 0.0, 0.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Negating Tuple {%v, %v, %v, %v} (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,

			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			testTup.a.Negate()
			if !TupleEquals(testTup.a, testTup.result) {
				t.Errorf("Cannot correctly Negate tuples using standalone function!")
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	var testTuples = []struct {
		a      Tuple
		scalar float64
		result Tuple
	}{
		{Tuple{1.0, -2.0, 3.0, -4.0}, 3.5, Tuple{3.5, -7.0, 10.5, -14.0}},
		{Tuple{1.0, -2.0, 3.0, -4.0}, 0.5, Tuple{0.5, -1.0, 1.5, -2.0}},
		{Tuple{0.0, 0.0, 0.0, 0.0}, 1.0, Tuple{0.0, 0.0, 0.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Multiplying Tuple {%v, %v, %v, %v} by Scalar %v (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.scalar,
			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			test := Multiply(testTup.a, testTup.scalar)
			if !TupleEquals(test, testTup.result) {
				t.Errorf("Cannot correctly Multiply tuples")
			}
		})
	}
}

func TestDivide(t *testing.T) {
	var testTuples = []struct {
		a       Tuple
		divisor float64
		result  Tuple
	}{
		{Tuple{1.0, -2.0, 3.0, -4.0}, 0.5, Tuple{2.0, -4.0, 6.0, -8.0}},
		{Tuple{1.0, -2.0, 3.0, -4.0}, 2.0, Tuple{0.5, -1.0, 1.5, -2.0}},
		{Tuple{0.0, 0.0, 0.0, 0.0}, 1.0, Tuple{0.0, 0.0, 0.0, 0.0}},
		{Tuple{1.0, -2.0, 3.0, -4.0}, 0.0, Tuple{0.0, 0.0, 0.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Dividing Tuple {%v, %v, %v, %v} by divisor %v (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.divisor,
			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			test := Divide(testTup.a, testTup.divisor)
			if !TupleEquals(test, testTup.result) {
				t.Errorf("Cannot correctly Divide tuples")
			}
		})
	}
}

func TestMagnitude(t *testing.T) {
	var testTuples = []struct {
		a      Tuple
		result float64
	}{
		{Tuple{1.0, 0.0, 0.0, 1.0}, 1.0},
		{Tuple{0.0, 1.0, 0.0, 1.0}, 1.0},
		{Tuple{0.0, 1.0, 1.0, 0.0}, 0.0},
		{Tuple{0.0, 0.0, 1.0, 1.0}, 1.0},
		{Tuple{1.0, 2.0, 3.0, 1.0}, math.Sqrt(14.0)},
		{Tuple{-1.0, -2.0, -3.0, 1.0}, math.Sqrt(14.0)},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Computing Magnitude for Tuple {%v, %v, %v, %v} (Expected Result %v)",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.result)
		t.Run(testName, func(t *testing.T) {
			mag := testTup.a.Magnitude()
			if !Equals(mag, testTup.result) {
				t.Errorf("Cannot correctly compute Magnitude for tuples using standalone function! Tuple: {%v,%v,%v,%v}; Calculated magnitude: %v; Expected Magnitude: %v",
					testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W, mag, testTup.result)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	var testTuples = []struct {
		a, result Tuple
	}{
		{Tuple{4.0, 0.0, 0.0, 1.0}, Tuple{1.0, 0.0, 0.0, 1.0}},
		{Tuple{1.0, 2.0, 3.0, 1.0}, Tuple{1.0 / math.Sqrt(14.0), 2.0 / math.Sqrt(14.0), 3.0 / math.Sqrt(14.0), 1.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Normalizing Tuple {%v, %v, %v, %v} (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.result.X, testTup.result.Y, testTup.result.Z, testTup.result.W)
		t.Run(testName, func(t *testing.T) {
			normed := testTup.a.Normalize()
			if !TupleEquals(normed, testTup.result) {
				t.Errorf("Cannot correctly Normalize tuples using standalone function! Tuple values are {%v, %v, %v, %v}",
					normed.X, normed.Y, normed.Z, normed.W)
			}
		})
	}
}

func TestDot(t *testing.T) {
	var testTuples = []struct {
		a, b   Tuple
		result float64
	}{
		{Tuple{1.0, 2.0, 3.0, 0.0}, Tuple{2.0, 3.0, 4.0, 0.0}, 20.0},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Dotting Tuple {%v, %v, %v, %v} with Tuple {%v, %v, %v, %v}, Expected Product: %v",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.b.X, testTup.b.Y, testTup.b.Z, testTup.b.W,
			testTup.result)
		t.Run(testName, func(t *testing.T) {
			dotProduct := Dot(testTup.a, testTup.b)
			if !Equals(dotProduct, testTup.result) {
				t.Errorf("Cannot correctly Dot tuples using standalone function! Dot product: %v",
					dotProduct)
			}
		})
	}
}

func TestCross(t *testing.T) {
	var testTuples = []struct {
		a, b, result Tuple
	}{
		{Tuple{1.0, 2.0, 3.0, 0.0}, Tuple{2.0, 3.0, 4.0, 0.0}, Tuple{-1.0, 2.0, -1.0, 0.0}},
		{Tuple{2.0, 3.0, 4.0, 0.0}, Tuple{1.0, 2.0, 3.0, 0.0}, Tuple{1.0, -2.0, 1.0, 0.0}},
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Crossing Tuple {%v, %v, %v, %v} with Tuple {%v, %v, %v, %v}",
			testTup.a.X, testTup.a.Y, testTup.a.Z, testTup.a.W,
			testTup.b.X, testTup.b.Y, testTup.b.Z, testTup.b.W)
		t.Run(testName, func(t *testing.T) {
			crossProduct := Cross(testTup.a, testTup.b)
			if !TupleEquals(crossProduct, testTup.result) {
				t.Errorf("Cannot correctly Dot tuples using standalone function! Cross product: Tuple{%v,%v,%v,%v}",
					crossProduct.X, crossProduct.Y, crossProduct.Z, crossProduct.W)
			}
		})
	}
}
