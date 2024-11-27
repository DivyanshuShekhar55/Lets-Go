package main

import "fmt"

func multiReturn(x, y int) (int, int) {
	return x * 2, y * 3
}

func multiReturnWithIgnore(x int) (int, int) {
	return x + 2, x + 3
}

func namedReturns(length, width int) (area, ratio int) {
	// named export vars initialised with a zero value like area = 0
	// good for documentation purposes, like user knows the return values are

	area = length * width
	ratio = length / width

	return area, ratio
}

func func_main() {

	a := 10
	fmt.Printf("hey this should come as a string : %d ", a)
	fmt.Println(a)
	x2, y3 := multiReturn(5, 8)
	fmt.Println(x2, y3)

	notIgnored, _ := multiReturnWithIgnore(10)
	// remember that Go cannot have unused variables, to ignore something you don't require use the '_'
	fmt.Println(notIgnored)

	a, b := namedReturns(10, 5)
	fmt.Println(a, b)

}
