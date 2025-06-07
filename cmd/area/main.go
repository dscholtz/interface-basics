package main

import (
	"fmt"

	"github.com/dscholtz/interface-basics.git/pkg/shape"
)

func main() {
	shapes := []shape.Shape{
		shape.Circle{Radius: 2.5},
		shape.Rectangle{Width: 3, Height: 4},
	}

	for _, s := range shapes {
		fmt.Printf("Area: %.2f\n", s.Area())
	}
}
