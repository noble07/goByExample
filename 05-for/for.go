package main

import "fmt"

func main() {

	i := 1
	// single condition for
	for i <= 3 {
		fmt.Println(i)
		i += 1 // or i = i + 1
	}

	// classic initial/condition/after for loop
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// using range
	for i := range 3 {
		fmt.Println("range", i)
	}

	// you can break out of the loop by using break or return from the enclosing function
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
