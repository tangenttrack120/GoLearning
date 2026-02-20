package main

import (
	"fmt"
)

func main() {
	// Go basic data types
	var a bool = true     // Boolean
  	var b int = 5         // Integer
  	var c float32 = 3.14  // Floating point number
  	var d string = "Hi!"  // String
	fmt.Println("Boolean: ", a, "\nInteger: ", b, "\nFloating point number: ", c, "\nString: ", d)

	// memory allocation for integer
	// int is a signed integer type that can hold both positive and negative values. The size of int is platform-dependent, but it is typically 32 or 64 bits.
	// int8 is a signed 8-bit integer type that can hold values from -128 to 127.
	// int16 is a signed 16-bit integer type that can hold values from -32,768 to 32,767.
	// int32 is a signed 32-bit integer type that can hold values from -2,147,483,648 to 2,147,483,647.
	// int64 is a signed 64-bit integer type that can hold values from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.
	var e int8 = -128
	var f int16 = -32768
	var g int32 = -2147483648
	var h int64 = -9223372036854775808
	fmt.Println("int8: ", e, "\nint16: ", f, "\nint32: ", g, "\nint64: ", h)

	// memory allocation for unsigned integer
	// uint is an unsigned integer
	// uint8 is an unsigned 8-bit integer type that can hold values from 0 to 255.
	// uint16 is an unsigned 16-bit integer type that can hold values from 0 to 65,535.
	// uint32 is an unsigned 32-bit integer type that can hold values from 0 to 4,294,967,295.
	// uint64 is an unsigned 64-bit integer type that can hold values from 0 to 18,446,744,073,709,551,615.
	var i uint8 = 255
	var j uint16 = 65535
	var k uint32 = 4294967295
	var l uint64 = 18446744073709551615
	fmt.Println("uint8: ", i, "\nuint16: ", j, "\nuint32: ", k, "\nuint64: ", l)
	
	// memory allocation for floating point numbers
	// float32 is a 32-bit floating-point number that can represent a wide range of values with a precision of about 7 decimal places.
	// float64 is a 64-bit floating-point number that can represent a wider range of values with a precision of about 15 decimal places.
	var m float32 = 3.14
	var n float64 = 3.141592653589793
	fmt.Println("float32: ", m, "\nfloat64: ", n)

	// arrays in Go are fixed-size collections of elements of the same type. The size of an array is determined at compile time and cannot be changed at runtime.
	// 2 ways to declare an array in Go
	var arr1 [5]int // declares an array of 5 integers
	arr2 := [3]string{"Hello", "World", "!"} // declares and initializes an array of 3 strings
	fmt.Println("Array 1: ", arr1, "\nArray 2: ", arr2)

	// we can also declare an array without specifying the size, and Go will infer the size based on the number of elements in the initializer
	arr3 := [...]float64{3.14, 2.718, 1.618} // declares and initializes an array of 3 floating-point numbers
	fmt.Println("Array 3: ", arr3)

	// we can access individual elements of an array using their index, which starts at 0
	fmt.Println("First element of arr2: ", arr2[0])
	fmt.Println("Second element of arr2: ", arr2[1])
	fmt.Println("Third element of arr2: ", arr2[2])

	// change the value of an element in an array by assigning a new value to it using its index
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	arr1[3] = 40
	arr1[4] = 50
	fmt.Println("Updated Array 1: ", arr1)

	// initialize an array with default values
	var arr4 [5]int // declares an array of 5 integers with default values of 0
	fmt.Println("Array 4: ", arr4)

	// we can also use the built-in function len() to get the length of an array
	fmt.Println("Length of arr1: ", len(arr1))
	fmt.Println("Length of arr2: ", len(arr2))
	fmt.Println("Length of arr3: ", len(arr3))
	fmt.Println("Length of arr4: ", len(arr4))
	
	// we can also use the built-in function copy() to copy the elements of one array to another
	var arr5 [5]int
	copy(arr5[:], arr1[:]) // copies the elements of arr1 to arr5
	fmt.Println("Array 5: ", arr5)

	// Initialize only specific elements of an array, and the rest will be initialized with default values
	arr6 := [5]int{0: 1, 2: 3} // initializes the first and third elements of the array, and the rest will be initialized with default values of 0
	fmt.Println("Array 6: ", arr6)

	// we can also use the built-in function append() to add elements to an array, but it returns a new slice with the added elements
	arr7 := []int{1, 2, 3} // declares and initializes a slice of 3 integers
	arr7 = append(arr7, 4) // adds the element 4 to the slice
	fmt.Println("Slice 7: ", arr7)

	// Go slices are more flexible than arrays, as they can grow and shrink dynamically. A slice is a reference to an underlying array, and it has a length and a capacity. The length of a slice is the number of elements it contains, while the capacity is the number of elements in the underlying array that can be accessed through the slice.
	// create slice using []dataType{values}
	slice1 := []int{1, 2, 3, 4, 5} // declares and initializes a slice of 5 integers
	fmt.Println("Slice 1: ", slice1)

	// two functions to return the length and capacity of a slice
	fmt.Println("Length of slice1: ", len(slice1)) // returns the length of the slice, which is the number of elements it contains
	fmt.Println("Capacity of slice1: ", cap(slice1)) // returns the capacity of the slice, which is the number of elements in the underlying array that can be accessed through the slice

	// create a slice from an array using the slicing syntax
	arr8 := [5]int{1, 2, 3, 4, 5} // declares and initializes an array of 5 integers
	slice2 := arr8[1:4] // creates a slice that references the elements from index 1 to index 3 of the array
	fmt.Println("Slice 2: ", slice2)

	// we can also create a slice that references the entire array using the slicing syntax
	slice3 := arr8[:] // creates a slice that references all the elements of the array
	fmt.Println("Slice 3: ", slice3)

	// we can also create a slice with a specific length and capacity using the built-in function make()
	// slice := make([]dataType, length, capacity)
	slice4 := make([]int, 5) // creates a slice of integers with length 5 and capacity 5
	fmt.Println("Slice 4: ", slice4)

	// access individual elements of a slice using their index, which starts at 0
	fmt.Println("First element of slice1: ", slice1[0])
	fmt.Println("Second element of slice1: ", slice1[1])
	fmt.Println("Third element of slice1: ", slice1[2])

	// change the value of an element in a slice by assigning a new value to it using its index
	slice1[0] = 10
	slice1[1] = 20
	slice1[2] = 30
	slice1[3] = 40
	slice1[4] = 50
	fmt.Println("Updated Slice 1: ", slice1)

	fmt.Println("Slice1: ", slice1)
	fmt.Println("Slice1 length: ", len(slice1))
	fmt.Println("Slice1 capacity: ", cap(slice1))

	// append element to slice
	// sliceName = append(SliceName, element1, element2, ...)
	slice1 = append(slice1, 60, 70)
	fmt.Println("Slice1 now:", slice1)
	fmt.Println("Slice1 length now: ", len(slice1))
	fmt.Println("Slice1 capacity now: ", cap(slice1))

	// append 1 slice to another slice
	sliceint1 := []int{1,2,3}
	sliceint2 := []int{4,5,6}
	sliceint3 := append(sliceint1, sliceint2...)
	fmt.Println("sliceint3: ", sliceint3)
	fmt.Println("sliceint3 length: ", len(sliceint3))
	fmt.Println("sliceint3 capacity: ", cap(sliceint3))

	// chage lenth of slice: unlike arrays it is possible to change lenght of an array
	sliceint3 = sliceint3[0:4] // change the lenght of slice to 4
	fmt.Println("sliceint3 now: ", sliceint3)
	fmt.Println("sliceint3 length now: ", len(sliceint3))
	fmt.Println("sliceint3 capacity now: ", cap(sliceint3))

	// Memory efficiency
	// if the array is large and you only need few elements it's better to use copy() function to copy it to slice
	// copy() function creates a new underlying array with teh required elements for the slice. This will reduce memory usage and improve performance
	// copy(dest, src)
	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15} // Original slice
  	fmt.Printf("numbers = %v\n", numbers)
  	fmt.Printf("length = %d\n", len(numbers))
  	fmt.Printf("capacity = %d\n", cap(numbers))

  	// Create copy with only needed numbers
  	neededNumbers := numbers[:len(numbers)-10]
  	numbersCopy := make([]int, len(neededNumbers))
  	copy(numbersCopy, neededNumbers)

  	fmt.Printf("numbersCopy = %v\n", numbersCopy)
  	fmt.Printf("length = %d\n", len(numbersCopy))
  	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}