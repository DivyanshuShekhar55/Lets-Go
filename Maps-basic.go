package main

import (
	//"errors"
	"fmt"
)

type Car struct {
	name      string
	top_speed int
}


// 1. Maps are just key:value pairs, like dicts in Python or maps in JS
// used for faster searching or storing unique key-values
// Creating A Map in Go ...
func createMap() {

	// syntax is make(map[type1]type2), type1 and type2 can be same as well
	// type 2 can be any type,
	// type1 can be only comparable values, so must not be maps, slices or functions
	myMap := make(map[string]string)
	myMap["Harry"] = "Potter"
	myMap["Peter"] = "Parker"

	// can create by direct initialisation also...
	ageMap := map[string]int{
		"Tortoise": 150,
		"Human":    100,
	}

	ageMap["Dog"] = 10
}


// **************************************************************************
// 2. Implemeting Basic Functionalities
func basicFunctions() {
	fmt.Println("\nInside The Basic Function Section: ")

	carMap := make(map[int]Car)

	carMap[1] = Car{
		name:      "Ferrari",
		top_speed: 340,
	}

	carMap[2] = Car{
		name:      "Lamborghini",
		top_speed: 350,
	}

	// find for a value
	car_at_rank2 := carMap[2]
	fmt.Println(car_at_rank2)

	// delete from map using the key
	delete(carMap, 2) // deletes the lamborghini entry

	// check if a key exists or not ...
	// if key doesn't exists returns me a 0 value...
	elem, ok := carMap[2]
	if !ok {
		fmt.Println("Value of element if key Not Found ❌:", elem) // returns 0
	} else {
		fmt.Println("Value of element if key exists ✅: ", elem) // returns the element
	}

}


// **************************************************************************
// 3. Nested Values in maps
func add(name_letter_map map[rune]map[string]int, names []string) {
	for i := 0; i < len(names); i++ {
		name := names[i]
		first_char := rune(name[0])
		_, ok := name_letter_map[first_char]
		if !ok {
			// the character was not present earlier so we need to ...
			// ... create a map corresponding to the value of this char (like 'A')
			mp := make(map[string]int)
			name_letter_map[first_char] = mp
		}
		name_letter_map[first_char][name]++
	}
}


func main() {
	basicFunctions()

	// example for point 3. nested map values ...
	name_letter_map := make(map[rune]map[string]int)
	// stores like {'A' : {'Antman': 5}}, like all 'A' names together and Antman showed up 5 times

	test_names := []string{"Antman", "Batman", "Antman", "Superman"}
	add(name_letter_map, test_names)
	// remember that both maps, slices will be passed by reference inside the function

	for key, val := range name_letter_map {
		fmt.Printf("%c: %v\n", key, val) // Use %c to print the character instead of its numeric ASCII value
	}

}
