package main

import (
	"fmt"
)

// Errors in Go are different from Error-handling in other langs like JS where we have try-catch
// In try blocks we might put many "dangerous" functions together, with just a single error logic
// In Go, we are forced to think about all types of errors separately
// All errors implement the 'error' interface


// ***********************************************************************
// 1. Returning Errors From Functions
func getUser(id int) (name string, err error) {
	if id < 0 {
		return "", fmt.Errorf("Invalid ID")
	}
	return "ABC", nil
}

// Handling Errors :
func greetUser(id int) (greeting string, err error) {
	user_name, err := getUser(id)

	if err != nil {
		return "", fmt.Errorf("Error occured: %s", err)
	} else {
		return fmt.Sprintf("Hey %s", user_name), nil
	}
}


// ************************************************************************
// 2. Error Interface and Handling Errors Like a Pro ...
type LoginError struct {
	msg      string
	code     int
	priority int
}

// use the in-built error interface
func (loginErr LoginError) Error() string {
	return fmt.Sprintf("%d: Error while Login. %s", loginErr.code, loginErr.msg)
}


func err_main() {
	greetings, err := greetUser(-55)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(greetings)
	}

	// Example for 2. Using the Error Interface
	// lets say we encounter error as not registered
	newLoginErr := LoginError{code: 404, msg: "Not Registered"}
	fmt.Println(newLoginErr.Error())

}
