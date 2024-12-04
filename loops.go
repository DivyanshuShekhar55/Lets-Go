package main

import "fmt"

// For loops - code pieces that run multiple times
func forLoop(n int) {
	fmt.Println("For loop Ahead ...")
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}

	// the scope of i has ended so we can re-initialise it
	// the following loop counts in the backward manner
	for i := n; i >= 0; i-- {
		fmt.Println(i)
	}
}

// ******************************************************************
// While Loops - in Go we don't have any separate "while" loop
// it is just a for loop written differently
func whileLoop(n int) {
	i := 0
	for i <= n {
		fmt.Println(i)
		i++
	}
}

func _main() {
	forLoop(5)

	fmt.Println("\nWhile Loop Ahead ...")
	whileLoop(5)
}
