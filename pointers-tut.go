package main

import (
	"fmt"
)

type car struct{
	model string
	price float32
}

func (car *car) updatePrice(new_price float32){
	car.price = new_price
	// methods are very often executed via pointers
	// this ensures that the value "really" updates (not a pass by value)
}

func updatePointerValue(ptr *int){
	*ptr = *ptr - 2
	// this is called passing by reference
	// updates value of ptr even though we haven't returned the new ptr value
}

func main() {

	// Pointers hold the address of some data value
	// very similar to C++, except that Pointer Arithmetic not allowed
	x := 5
	ptr := &x // same as: var ptr *int = &x
	fmt.Println("value of ptr:", ptr, "\nvalue of *ptr:", *ptr)

	*ptr = (*ptr) + 5
	fmt.Println("val of 'x' after updating *ptr value:", x)

	updatePointerValue(ptr)
	fmt.Println("*ptr after passing to function:", *ptr)

	batmobile := car{model: "Batmobile", price: 10000.00}
	batmobile.updatePrice(99999.99) // notice that we didn't pass pointer, auto type-casting occurs
	fmt.Println(batmobile.price) 
}
