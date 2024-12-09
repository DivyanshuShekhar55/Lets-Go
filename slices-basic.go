package main

import (
	"fmt"
)

// Arrays in Go are of fixed size, just like in C++
func count(person1 string, person2 string, person3 string) int {

	// if not initialised the defualt values are given to array like 0 or ""
	names := [3]string{person1, person2, person3}
	// fmt.Println(names)

	// follows 0-based indexing
	// fmt.Println(names[1])

	return len(names)

}


// ******************************************************************************
// Slices In Go
// Slices are resizable arrays, just like vectors in C++
// Slices always 'refer' to some kind of array, so their max capacity to 'resize' is length of array
// However, we need not always explicitly create an array to create a slice
// if we need to grow a slice beyond its capacity, we need to copy the contents to a new larger location


// *******************************************************************************
// Creating Slices with Arrays
func experimentSlices() {

	greet := [5]string{"Hello", "Namaste", "Konnichiwa", "Hola", "Bonjour"}

	slice_one := greet[0:3]  // create a slice from index-0 (including) to 3 (excluding)
	slice_two := greet[:3]   // from staring (index-0) till index-3 (excluding index-3)
	slice_three := greet[3:] // from index-3 (inluding) till last (including)
	slice_four := greet[:]   // complete array, start to end

	fmt.Println(slice_one, "\n", slice_two, "\n", slice_three, "\n", slice_four)

}

func usingSlices() {
	// We can make slices without arrays too
	s := make([]int, 5, 8) // int slice of size 5 and capacity 8
	// if capacity is not provided then it defaults to the length
	s = append(s, 5) // add new element in th list at the last unused index

}


// *********************************************************************************
// Variadic function = accepts any number of arguments, like Println function
func variadic(nums ...int) {
	// any number of integer vars will be accepted
	// will be passed in as a slice
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}


// Spread Operator = one slice passed in, you take out everything like individual items
// pass a slice to a variadic function
func spreadExample() {
	primes := []int{1, 3, 5, 7}
	fmt.Println("\nPrime Numbers ")
	variadic(primes...) // the ... after primes is spread operator
}


// **********************************************************************************
// adding new elements to slices
func adding() {
	names := []string{"Mickey", "Donald", "Pooh"}

	// ✅ always re-assign the returned new slice
	// ❌ DON'T : someSlice = append(slice, newItem)
	names = append(names, "Superman")
}

func main() {
	fmt.Println("Length is", count("Harry", "PrimeAgen", "FreecodeCamp"))
	experimentSlices()

	variadic(1, 2, 3, 3, 4) // call variadic function
	spreadExample()
}
