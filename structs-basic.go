package main

import "fmt"

type user struct {
	name string
	id   int64
}

type msg struct {
	// nested struct, struct within struct
	text     string
	sender   user
	receiver user
}

type admin struct {
	// embedded struct, kind of similar to inheritance
	// all the fields of user will be "embedded" here also
	user
	department string
}

func struct_main() {
	john := user{} //empty struct
	john.id = 12345
	john.name = "John Doe"
	fmt.Println(john)

	harry := admin{
		user: user{
			name: "Harry Potter",
			id:   123456,
		},
		department: "detective",
	}

	fmt.Println(harry)
	// not nested as we can access fields directly as shown ...
	fmt.Println(harry.name)
	

}
