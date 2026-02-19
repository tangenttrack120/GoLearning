// GO Syntax
// consist of package declaration
package main

// import statement
import (
	"fmt"
)

// main function
func main() {
	// best practice for naming variables in Gois to use camelCase
	// to declare a variable, we use the var keyword followed by the variable name and type
	var firstName string = "Richard"
	// we can also declare a variable without initializing it, and assign a value to it later
	var lastName string
	lastName = "Victor"
	// we can also declare a variable without specifying the type, and Go will infer the type based on the value assigned to it
	age := 30
	// we can also declare multiple variables at once
	var x, y, z int = 1, 2, 3
	a, b, c := 4,5,6

    fmt.Println("Hello", firstName, lastName)
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	x = x + y + z
	a = a + b + c
}
