package main

import (
	"fmt"

	"github.com/jleumas/ray/tuples"
)

func main() {
	projectile := Projectile{
		tuples.Point(0.0, 1.0, 0.0),
		tuples.Vector(1.0, 1.0, 0.0),
	}

	environment := Environment{
		tuples.Vector(0.0, -0.1, 0.0),
		tuples.Vector(-0.01, 0.0, 0.0),
	}

	for x := 0; x < 100; x++ {
		projectile = tick(environment, projectile)
		fmt.Println(projectile)
	}

}

type Projectile struct {
	point    tuples.Tuple
	velocity tuples.Tuple
}

type Environment struct {
	gravity tuples.Tuple
	wind    tuples.Tuple
}

func tick(e Environment, p Projectile) Projectile {
	p.point = tuples.Add(p.point, p.velocity)
	p.velocity = tuples.Add(p.velocity, e.gravity, e.wind)
	return Projectile{
		p.point,
		p.velocity,
	}
}
