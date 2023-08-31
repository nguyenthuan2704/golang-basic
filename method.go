package main

import (
	"fmt"
	"math"
)

type Vetex struct {
	X, Y float64
}

func (v Vetex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vetex{3, 4}
	fmt.Println(v.Abs())
}
