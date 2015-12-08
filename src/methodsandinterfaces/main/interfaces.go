package main

import (
	"math"
	"fmt"
)

type Abser interface {
	Absolute() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	fmt.Println(a.Absolute())
}

type MyFloat float64

func (f MyFloat) Absolute() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Absolute() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


