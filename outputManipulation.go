package main

import (
	"fmt"
)

func main() {
	// Constants in Go are declared using the const keyword, and their value cannot be changed after they are assigned
	const pi float64 = 3.14
	const gravity = 9.8

	// We can also declare multiple constants at once
	const (
		lightSpeed = 299792458 // in meters per second
		planckConstant = 6.62607015e-34 // in joule seconds
	)
	const planetName, countryName, cityName = "Earth", "USA", "New York"

	// have 3 function to print output
	// fmt.Println() prints the output with a newline at the end
	fmt.Println("Hello, World!")
	// fmt.Printf() prints the output with a formatted string
	fmt.Printf("The value of pi is: %v\n", pi)
	// fmt.Print() prints the output without a newline at the end
	fmt.Print("The value of gravity is: ", gravity)

	// general Formating Verbs
	// %v - default format
	// %T - type of the variable
	// %d - decimal integer
	// %f - floating-point number
	// %s - string
	// %t - boolean
	fmt.Printf("The value of lightSpeed is: %v m/s\n", lightSpeed)
	fmt.Printf("The value of planckConstant is: %v J*s\n", planckConstant)
	fmt.Printf("The planet name is: %s\n", planetName)
	fmt.Printf("The country name is: %s\n", countryName)
	fmt.Printf("The city name is: %s\n", cityName)

	// Integer Formating Verbs
	// %b - binary
	// %o - octal
	// %x - hexadecimal
	// %X - hexadecimal with uppercase letters
	fmt.Printf("The value of lightSpeed in binary is: %b\n", lightSpeed)
	fmt.Printf("The value of lightSpeed in octal is: %o\n", lightSpeed)
	fmt.Printf("The value of lightSpeed in hexadecimal is: %x\n", lightSpeed)
	fmt.Printf("The value of lightSpeed in hexadecimal with uppercase letters is: %X\n", lightSpeed)

	// we can also specify the width and precision of the output using the format specifiers
	// %4d - minimum width of 4 characters for a decimal integer
	// %-4d - left-align the output within the specified width
	// %04d - pad the output with zeros to fill the specified width
	i := 42
	fmt.Printf("%b\n", i)
  	fmt.Printf("%d\n", i)
  	fmt.Printf("%+d\n", i)
  	fmt.Printf("%o\n", i)
  	fmt.Printf("%O\n", i)
  	fmt.Printf("%x\n", i)
  	fmt.Printf("%X\n", i)
  	fmt.Printf("%#x\n", i)
  	fmt.Printf("%4d\n", i)
  	fmt.Printf("%-4d\n", i)
  	fmt.Printf("%04d\n", i)

	// Floating-point Formating Verbs
	// %e - scientific notation
	// %E - scientific notation with uppercase E
	// %f - decimal point but no exponent
	// %.2f - decimal point with 2 digits of precision
	// %6.2f - minimum width of 6 characters with 2 digits of precision	
	// %g - uses the shortest representation of %e or %f
	fmt.Printf("The value of planckConstant in scientific notation is: %e\n", planckConstant)
	fmt.Printf("The value of planckConstant in scientific notation with uppercase E is: %E\n", planckConstant)
	fmt.Printf("The value of planckConstant in decimal point but no exponent is: %f\n", planckConstant)
	fmt.Printf("%.2f\n", i)
  	fmt.Printf("%6.2f\n", i)
	fmt.Printf("The value of planckConstant in shortest representation is: %g\n", planckConstant)

	// boolean Formating Verbs
	// %t - boolean value
	isRaining := true
	fmt.Printf("Is it raining? %t\n", isRaining)

	// String Formating Verbs
	// %s - string
	// %8s	Prints the value as plain string (width 8, right justified)
	// %-8s	Prints the value as plain string (width 8, left justified)
	// %q - double-quoted string with Go syntax escaping
	// %x - hexadecimal representation of the string
	// %X - hexadecimal representation of the string with uppercase letters
	name := "Go Programming"
	fmt.Printf("The name of the language is: %s\n", name)
	fmt.Printf("The name of the language in 8-character width right justified is: %8s\n", name)
	fmt.Printf("The name of the language in 8-character width left justified is: %-8s\n", name)
	fmt.Printf("The name of the language in double-quoted string with Go syntax escaping is: %q\n", name)
	fmt.Printf("The name of the language in hexadecimal representation is: %x\n", name)
	fmt.Printf("The name of the language in hexadecimal representation with uppercase letters is: %X\n", name)
}