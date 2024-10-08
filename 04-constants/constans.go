package main

import (
	"fmt"
	"math"
)

// const declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	// const statement can appear anywhere a var statement can
	const n = 500000000

	// constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
