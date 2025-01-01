package main

import (
	"fmt"
)

// Note that these functions takes only two inputs (or operands) ...

func sum(a, b int) (res int) {
	return a + b
}

func diff(a, b int) (res int) {
	return a - b
}

// ****************************************************************************
// Higher Order Functions
// Functions that take other functions as input
func aritmeticOfThree(a, b, c int, arithmetic func(int, int) int) (res int) {

	// basically takes 3 nums and a func, returns res as response
	// the input function is called First Class Func. : arithmetic func(int, int) int

	// arithmetic functions takes two input, returns an int as a result
	first_res := arithmetic(a, b)
	final_res := arithmetic(first_res, c)
	return final_res
}

// ****************************************************************************
// Function Currying
// Higher Order Func., whose return value is also a Function
func aritmeticOfThree_Curry(arithmetic func(int, int) int) func(int, int, int) int {

	// think in reverse way - "I have to return this type of func using this type of input func."
	// assume var names yourself ...
	return func(a, b, c int) int {
		first_res := arithmetic(a, b)
		final_res := arithmetic(first_res, c)
		return final_res
	}

}

// ******************************************************************************
// Closures

func advanced_function_main() {
	fmt.Println(aritmeticOfThree(3, 4, 5, sum))
	fmt.Println(aritmeticOfThree(10, 5, 0, diff))

	println("\n")
	sumOfThree := aritmeticOfThree_Curry(sum) // this is a function now
	fmt.Println(sumOfThree(3, 4, 5))

	diffOfThree := aritmeticOfThree_Curry(diff) // also a function
	fmt.Println(diffOfThree(10, 5, 0))


	// **********************************************************************
	// Immediately Invoked Function Expressions (IIFE)
	func() {
		// Anonymous function, which is called immediately after its body finishes
		// we don't need to call this function explicitly, also can be called only once
		// can be used for tasks like DB connection
		fmt.Println("The IIFE was called")
	}()
}
