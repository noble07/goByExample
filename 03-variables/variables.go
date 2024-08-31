package main

import "fmt"

func main() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// go infers type of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued
	var e int
	fmt.Println(e)

	// the := is a syntax shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)
}
