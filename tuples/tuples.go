/*
Package tuples provides low-level data primitives and operations required for use in a ray tracing library
*/
package tuples

import (
	"math"
)

func Equals(a, b float64) bool {
	const Epsilon = 0.00001
	if math.Abs(a-b) < Epsilon {
		return true
	} else {
		return false
	}
}

//Type Tuple provides the basic three-dimensional scalars
//and a field that designates whether the tuple is a point or a vector.
type Tuple struct {
	X, Y, Z, W float64
}

func (t Tuple) IsPoint() bool {
	if t.W == 1.0 {
		return true
	} else {
		return false
	}
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func TupleEquals(a Tuple, b Tuple) bool {
	if Equals(a.X, b.X) &&
		Equals(a.Y, b.Y) &&
		Equals(a.Z, b.Z) &&
		Equals(a.W, b.W) {
		return true
	} else {
		return false
	}
}

func Add(nums ...Tuple) Tuple {
	newTup := Tuple{0.0, 0.0, 0.0, 0.0}
	for _, num := range nums {
		newTup.X += num.X
		newTup.Y += num.Y
		newTup.Z += num.Z
		newTup.W += num.W
	}

	return newTup
}

func Subtract(a, b Tuple) Tuple {
	newTup := Tuple{}
	newTup.X = a.X - b.X
	newTup.Y = a.Y - b.Y
	newTup.Z = a.Z - b.Z
	newTup.W = a.W - b.W

	return newTup
}

func (t *Tuple) Negate() {
	t.X *= -1
	t.Y *= -1
	t.Z *= -1
	t.W *= -1
}

func (t *Tuple) Multiply(scalar float64) {
	t.X *= scalar
	t.Y *= scalar
	t.Z *= scalar
	t.W *= scalar
}

func (t *Tuple) Divide(divisor float64) {
	if divisor == 0.0 {
		t.X = 0.0
		t.Y = 0.0
		t.Z = 0.0
		t.W = 0.0
	} else {
		t.X /= divisor
		t.Y /= divisor
		t.Z /= divisor
		t.W /= divisor
	}
}

func (t *Tuple) Magnitude() float64 {
	mag := math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
	return mag
}

func (t *Tuple) Normalize() Tuple {
	norm := Tuple{}
	mag := t.Magnitude()

	norm.X = t.X / mag
	norm.Y = t.Y / mag
	norm.Z = t.Z / mag
	norm.W = t.W

	return norm
}

func Dot(a, b Tuple) float64 {
	dot := a.X*b.X +
		a.Y*b.Y +
		a.Z*b.Z +
		a.W*b.W

	return dot
}

func Cross(a, b Tuple) Tuple {
	cross := Tuple{}

	cross.X = a.Y*b.Z - a.Z*b.Y
	cross.Y = a.Z*b.X - a.X*b.Z
	cross.Z = a.X*b.Y - a.Y*b.X
	cross.W = 0.0

	return cross
}
