package main

import "fmt"

func main() {
	// conditioning statements
	x := 20
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}

	// nested if statements
	y := 15
	if y > 10 {
		if y%2 == 0 {
			fmt.Println("y is greater than 10 and even")
		} else {
			fmt.Println("y is greater than 10 and odd")
		}
	} else if y == 10 {
		fmt.Println("y is equal to 10")
	} else {
		fmt.Println("y is less than 10")
	}

	// switch statement
	day := 2
	switch day {
	case 1 :
		fmt.Println("Monday")
	case 2 : 
		fmt.Println("Tuesday")
	default :
		fmt.Println("Invalid day")
	}

	// multi case switch statement
	switch day {
	case 1,3,5:
		fmt.Println("Odd weekday")
	case 2,4:
		fmt.Println("Even weekday")
	default:
		fmt.Println("Invalid day")
	}

	// looping statement
	// go only has for loop which can be used as while and do while loop
	for i := 0; i<5; i++ {
		// continue statement
		if i == 2 {
			fmt.Println("Skipping iteration")
			continue
		} else if i == 4 {
			// break statement
			fmt.Println("Breaking loop")
			break
		} else {
			 // nested loop
			for j := 0; j < 3; j++ {
				fmt.Println("i:", i, "j:", j)
			}
		}
	}

	// Range keyword: used to easily iterate through elements of array, slice or map. it returns both index and values
	fruits := [3]string{"apple", "orange", "banana"}
	for index, value := range fruits {
		fmt.Printf("index: %v\t value: %v\n", index, value)
	}

	// to only show the index or value use _
	for _, val := range fruits {
		fmt.Printf("value: %v", val)
	}
}