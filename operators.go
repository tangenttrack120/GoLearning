package main

import "fmt"

func main() {
	var x, y int = 10, 5

	// 1. Arithmetic Operators
	var (
		Addition       = x + y // 10 + 5 = 15
		Subtraction    = x - y // 10 - 5 = 5
		Multiplication = x * y // 10 * 5 = 50
		Division       = x / y // 10 / 5 = 2
		Modulus        = x % y // 10 % 5 = 0
	)

	// Increment and Decrement must be executed as standalone statements
	x++ // x becomes 11
	y-- // y becomes 4

	fmt.Println("Addition operator: ", Addition)
	fmt.Println("Subtraction operator: ", Subtraction) // Fixed typo here
	fmt.Println("Multiplication operator: ", Multiplication)
	fmt.Println("Division Operator: ", Division)
	fmt.Println("Modulus Operator: ", Modulus)
	fmt.Println("Increment Operator (x): ", x)
	fmt.Println("Decrement Operator (y): ", y)

	// Reset x and y for the rest of the examples
	x, y = 10, 5

	// 2. Assignment Operators
	var a = 2

	// We must declare variables with an initial value before using assignment operators
	var AddAssign, SubAssign, MulAssign, DivAssign, ModAssign = 0, 0, 1, 10, 10
	var NNAssign, OrAssign, XorAssign, RightShiftAssign, LeftShiftAssign = 10, 10, 10, 10, 10

	AddAssign += a        // 0 + 2 = 2
	SubAssign -= a        // 0 - 2 = -2
	MulAssign *= a        // 1 * 2 = 2
	DivAssign /= a        // 10 / 2 = 5
	ModAssign %= a        // 10 % 2 = 0
	NNAssign &= a         // 10 & 2 = 2
	OrAssign |= a         // 10 | 2 = 10
	XorAssign ^= a        // 10 ^ 2 = 8
	RightShiftAssign >>= a // 10 >> 2 = 2
	LeftShiftAssign <<= a  // 10 << 2 = 40

	fmt.Println("\nAddition Assignment Operator: ", AddAssign)
	fmt.Println("Subtraction Assignment Operator: ", SubAssign)
	fmt.Println("Multiplication Assignment Operator: ", MulAssign)
	fmt.Println("Division Assignment Operator: ", DivAssign)
	fmt.Println("Modulus Assignment Operator: ", ModAssign)
	fmt.Println("Bitwise AND Assignment Operator: ", NNAssign)
	fmt.Println("Bitwise OR Assignment Operator: ", OrAssign)
	fmt.Println("Bitwise XOR Assignment Operator: ", XorAssign)
	fmt.Println("Right Shift Assignment Operator: ", RightShiftAssign)
	fmt.Println("Left Shift Assignment Operator: ", LeftShiftAssign)

	// 3. Comparison Operators
	// Changed {} to () for variable grouping
	var (
		EqualTo            = x == y // 10 == 5 = false
		NotEqualTo         = x != y // 10 != 5 = true
		GreaterThan        = x > y  // 10 > 5 = true
		LessThan           = x < y  // 10 < 5 = false
		GreaterThanEqualTo = x >= y // 10 >= 5 = true
		LessThanEqualTo    = x <= y // 10 <= 5 = false
	)
	fmt.Println("\nEqual To Operator: ", EqualTo)
	fmt.Println("Not Equal To Operator: ", NotEqualTo)
	fmt.Println("Greater Than Operator: ", GreaterThan)
	fmt.Println("Less Than Operator: ", LessThan)
	fmt.Println("Greater Than Equal To Operator: ", GreaterThanEqualTo)
	fmt.Println("Less Than Equal To Operator: ", LessThanEqualTo)

	// 4. Logical Operators
	var (
		NNLogical  = (x == 10 && x > y)  // true
		OrLogical  = (x == 10 || x == 5) // Fixed capital 'X' to lowercase 'x'
		NotLogical = !(x == 10 || y == 5) // false
	)
	fmt.Println("\nAnd And Logical operator: ", NNLogical)
	fmt.Println("Or Logical Operator: ", OrLogical)
	fmt.Println("Not Logical Operator: ", NotLogical)

	// 5. Bitwise Operators
	var (
		AND        = x & y  
		OR         = x | y  
		XOR        = x ^ y  
		LeftShift  = x << y 
		RightShift = x >> y 
	)
	fmt.Println("\nAND Bitwise Operator: ", AND)
	fmt.Println("OR Bitwise Operator: ", OR)
	fmt.Println("XOR Bitwise Operator: ", XOR)
	fmt.Println("Zero Fill Left Shift Operator: ", LeftShift)
	fmt.Println("Zero Fill Right Shift Operator: ", RightShift) // Fixed typo here
}