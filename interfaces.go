package main

import (
	"fmt"
)

// 1. Interfaces are used to define method signatures
// every type can have own implementation of an interface

type organism interface {
	sound() string
}

type Human struct {
	// in Go, interfaces are implicit
	// we don't have to tell like type Human struct implements organism
	// so Human struct automatically is using the interface organism
	name string
}

// define how the human struct implements the sound func.
func (human Human) sound() string {
	return "Speech"
}

// ***********************************************************************
// 2. By naming our arguments and return types, we can increase readability
type play interface {
	game(age int, gender string) (fav_game string)
}

func (human Human) game(age int, gender string) string {
	if age <= 12 && gender == "male" {
		return fmt.Sprintf("Cars")
	}
	return fmt.Sprintf("can be anything")
}

// ************************************************************************
// 3. Type Assertions - assert / check if a var is of a particular type or not?
// we create an empty interface for this, see main function

//**************************************************************************
// 4. More complex type assertions

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}


// ***************************************************************************
// 5. Type Switches -- easier way to type assertions ...
func checkShape(shape interface{}) float64 {
	switch val := shape.(type) {
	case circle:
		return val.radius
	case float64:
		return val
	default:
		return -1
	}
}

func interfaces_main() {
	alex := Human{name: "Alex"}
	fmt.Println(alex.sound())

	// Example for Point 3. type assertions...
	var str interface{}
	str = "a"                        // this may happen anywhere in 10K lines of code
	returnVal, is_ok := str.(string) // is the value really string ?
	fmt.Println(returnVal, is_ok)

	// Example for 4. Complex type assertions
	var s shape              // shape is an interface
	s = circle{radius: 6.28} // giving interface a value of type circle
	i, ok := s.(circle)
	fmt.Println(i, ok)

}
