package main

import (
	"fmt"

	"github.com/jleumas/ray/tuples"
)

func main() {
	point := tuples.Point(1.0, 1.0, 1.0)
	if point == tuples.Point(1.0, 1.0, 1.0) {
		fmt.Println("Looks good")
	} else {
		fmt.Println("Looks bad")
	}
}
