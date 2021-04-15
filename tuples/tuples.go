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

type Tuple struct {
	x, y, z, w float64
}

func Point(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1.0}
}

func Vector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0.0}
}

func TupleEquals(a Tuple, b Tuple) bool {
	if Equals(a.x, b.x) &&
		Equals(a.y, b.y) &&
		Equals(a.z, b.z) &&
		Equals(a.w, b.w) {
		return true
	} else {
		return false
	}
}

func Add(a, b Tuple) Tuple {
	newTup := Tuple{}
	newTup.x = a.x + b.x
	newTup.y = a.y + b.y
	newTup.z = a.z + b.z

	if a.w == 1.0 && b.w == 1.0 {
		newTup.w = 1.0
	} else {
		newTup.w = a.w + b.w
	}

	return newTup
}

func Subtract(a, b Tuple) Tuple {
	newTup := Tuple{}
	newTup.x = a.x - b.x
	newTup.y = a.y - b.y
	newTup.z = a.z - b.z
	newTup.w = a.w - b.w

	return newTup
}

func (t *Tuple) Negate() {
	t.x *= -1
	t.y *= -1
	t.z *= -1
	t.w *= -1
}

func (t *Tuple) Multiply(scalar float64) {
	t.x *= scalar
	t.y *= scalar
	t.z *= scalar
	t.w *= scalar
}

func (t *Tuple) Divide(scalar float64) {
	t.x /= scalar
	t.y /= scalar
	t.z /= scalar
	t.w /= scalar
}

func (t *Tuple) Magnitude() float64 {
	mag := math.Sqrt(math.Pow(t.x, 2) + math.Pow(t.y, 2) + math.Pow(t.z, 2) + math.Pow(t.w, 2))
	return mag
}

func (t *Tuple) Normalize() Tuple {
	norm := Tuple{}
	mag := t.Magnitude()

	norm.x = t.x / mag
	norm.y = t.y / mag
	norm.z = t.z / mag
	norm.w = t.w / mag

	return norm
}

func Dot(a, b Tuple) float64 {
	dot := a.x*b.x +
		a.y*b.y +
		a.z*b.z +
		a.w*b.w

	return dot
}

func Cross(a, b Tuple) Tuple {
	cross := Tuple{}

	cross.x = a.y*b.z - a.z*b.y
	cross.y = a.z*b.x - a.x*b.z
	cross.z = a.x*b.y - a.y*b.x
	cross.w = 0.0

	return cross
}
