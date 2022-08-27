package main

import (
	"fmt"
	// Atribui um alias ao import
	m "math"
)

func main() {
	const PI float64 = 3.1415

	// tipo (float64) inferido pelo compilador
	var raio = 3.2

	// Forma reduzida de criar uma var
	area := PI * m.Pow(raio, 2)

	fmt.Println("A área da circunferência é", area)

	// blocos de construção consts
	const (
		a = 1
		b = 2
	)

	// blocos de construção vars
	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false

	fmt.Println(e, f)

	g, h, i := 2, false, "epa!"

	fmt.Println(g, h, i)
}
