package tuples

import (
	"fmt"
	"testing"
)

func TestPoint(t *testing.T) {
	var testPoints = []struct {
		x, y, z float64
	}{
		{0.0, 0.0, 0.0},
		{4.3, -4.2, 3.1},
		{-4.3, 4.2, -3.1},
	}

	for _, testPt := range testPoints {
		testName := fmt.Sprintf("%v, %v, %v", testPt.x, testPt.y, testPt.z)
		t.Run(testName, func(t *testing.T) {
			point := Point(testPt.x, testPt.y, testPt.z)
			if point.w != 1.0 {
				t.Errorf("Wrong w")
			}
			if point.x != testPt.x {
				t.Errorf("Wrong x")
			}
			if point.y != testPt.y {
				t.Errorf("Wrong y")
			}
			if point.z != testPt.z {
				t.Errorf("Wrong z")
			}
		})
	}
}

func TestVector(t *testing.T) {
	var testVectors = []struct {
		x, y, z float64
	}{
		{0.0, 0.0, 0.0},
		{4.3, -4.2, 3.1},
		{-4.3, 4.2, -3.1},
	}

	for _, testVec := range testVectors {
		testName := fmt.Sprintf("%v, %v, %v", testVec.x, testVec.y, testVec.z)
		t.Run(testName, func(t *testing.T) {
			point := Vector(testVec.x, testVec.y, testVec.z)
			if point.w != 0.0 {
				t.Errorf("Wrong w")
			}
			if point.x != testVec.x {
				t.Errorf("Wrong x")
			}
			if point.y != testVec.y {
				t.Errorf("Wrong y")
			}
			if point.z != testVec.z {
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
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Comparing Tuple {%v, %v, %v, %v} to Tuple {%v, %v, %v, %v} for equality (Expected Result: %v)",
			testTup.a.x, testTup.a.y, testTup.a.z, testTup.a.w,
			testTup.b.x, testTup.b.y, testTup.b.z, testTup.b.w,
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
			testTup.a.x, testTup.a.y, testTup.a.z, testTup.a.w,
			testTup.b.x, testTup.b.y, testTup.b.z, testTup.b.w,
			testTup.result.x, testTup.result.y, testTup.result.z, testTup.result.w)
		t.Run(testName, func(t *testing.T) {
			if !(TupleEquals(Add(testTup.a, testTup.b), testTup.result)) {
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
			testTup.a.x, testTup.a.y, testTup.a.z, testTup.a.w,
			testTup.b.x, testTup.b.y, testTup.b.z, testTup.b.w,
			testTup.result.x, testTup.result.y, testTup.result.z, testTup.result.w)
		t.Run(testName, func(t *testing.T) {
			if !(TupleEquals(Subtract(testTup.a, testTup.b), testTup.result)) {
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
	}

	for _, testTup := range testTuples {
		testName := fmt.Sprintf("Negating Tuple {%v, %v, %v, %v} (Expected Result Tuple {%v, %v, %v, %v})",
			testTup.a.x, testTup.a.y, testTup.a.z, testTup.a.w,

			testTup.result.x, testTup.result.y, testTup.result.z, testTup.result.w)
		t.Run(testName, func(t *testing.T) {
			if TupleEquals(testTup.a, testTup.result) {
				t.Errorf("Cannot correctly subtract tuples using standalone function!")
			}
		})
	}
}

func TestMultiply(t *testing.T) {

}

func TestDivide(t *testing.T) {

}

func TestMagnitude(t *testing.T) {

}

func TestDot(t *testing.T) {

}

func TestCross(t *testing.T) {

}
