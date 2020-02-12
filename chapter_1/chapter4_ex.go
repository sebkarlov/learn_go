package main

import (
	"fmt"
)

// declare an integer array of five elements
var array [5]int

// declare an integer array of five elements
// initialize each element with a specific value
array := [5]int{10, 20, 30, 40, 50}

// declare an integer array
// initialize each element with a specific value
// capacity is determined based on the number of values initialized
array := [...]int{10, 20, 30, 40, 50}

// declare an integer array of five elements
// initialized index 1 and 2 with speciefic values
// the rest of the elements contain their zero value
array := [5]int{1: 10, 2: 20}

// declare an integer array of five elements
// initialize each element with specific value
array := [5]int{10, 20, 30, 40, 50}
//change the value at index 2
array[2] = 35

// declare an integer pointer array of five elements
// initialize index 0 and 1 of the array with integer pointers
array := [5]int{0: new(int), 1:new(int)}
// assign values to index 0 and 1
*array[0] = 10
*array[1] = 20

// declare a string array of five elements
var array1 [5]string
// declare second array of five elements
// initialized the array with colors
array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
// copy the values from array2 into array1
array1 = array2

// declare a string array of four elements
var array1 [4]string
// declare a second string array of five elements
array2 := [5]string{"red", "blue", "green", "yellow", "pink"}
// copy the values from array2 into array1
// Compile error: cannot use array2 (type [5]string as type [4]string in assigment)

// declare a string pointer array of three elements
var array [3]*string
// declare a second string pointer array of three elements
// initialize the array with string pointers
array2 := [3]*string{new(string), new(string), new(string)}
// add colors to each element
*array2[0] = "red"
*array2[1] = "blue"
*array2[2] = "green"
// copy the values from array2 into array1
array1 = array2

// declare a two dimensional integer array of four elements by two elements
var array [4][2]int
// use array literal to declare and initialize a two dimensional integer array
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
// declare and initialize index 1 and 3 of the outer array
array := [4][2]int{1: {20, 21}, 3: {40, 41}}
// declare and initialize individual elements of the outer and inner array
array := [4][2]int{1: {0: 20}, 3: {1: 41}

// declare a two dimensional integer array of two elements
var array [2][2]int
// set integer values to each individual element
array[0][0] = 10
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40

// declare two diffrent two dimensional integer arrays
var array1 [2][2]int
var array2 [2][2]int
// add integer values to each individual element
array2[0][0] = 10
array2[0][1] = 20
array2[1][0] = 30
array2[1][1] = 40
// copy the values from array2 into array1
array1 = array2
// copy index 1 of array1 into a new array of the same type
var array3 [2]int = array1[1]
// copy the integer found in index 1 of outer array and index 0
// of the interior array into a new variable of type integer
var value int = array1[1][0]

// declare an array of 8 megabytes
var array [1e6]int
// pass the array to the function foo
foo(array)
// function foo accepts an array of one million integers
func foo(array [1e6]int) {
	...
}

// allocate an array of 8mb
var array [1e6]int
// pass the address of the array to the function foo
foo(&array)
// functoin foo accepts a pointer to an array of one million integers
func foo(array *[1e6]int) {
	...
}

//slices

// create a slice of strings
// contains a lenght and capacity of 5 elements
slice := make([]string, 5)

// create a slice of integers
// contains a lenght of 3 and has a capacity of 5 elements
slice := make([]int, 3, 5)

// create a slice of integers
// make the length larger than the capacity
slice := make([]int, 5, 3)
// Error, len larger than cap in make([]itn)

// create a slice of strings
// contains a lenght and capacity of 5 elements
slice := []string{"red", "blue", "green", "yellow", "pink"}
// create slice of integers lng and cap of 3 elements
slice := []int{10, 20, 30}

// create a slice of strings
// initialize the 100th element with an empty string
slice := []string{99: ""}

// diffrence between value inside []
// this is an array of 3 elements
array := [3]int{10, 20, 30}
// this is a slice with lenght and cap of 3
slice := []int{10, 20, 30}

// create a nil slice of integers
var slice []int

// use make to create an empty slice of integers
slice := make([]int, 0)
// use a slice literal to create am empty slice of integers
slice := []int{}

// create a slice of integers contains a len and cap of 5
slice := []int{10, 20, 30, 40, 50}
// change the value of index 1
slice[1] = 25

// create a slice of int len and cap of 5
slice := []int{10, 20, 30, 40, 50}
// create a new slice contains a len of 2 and cap of 4
newSlice := slice[1:3]
// for slice[i:j] with underlying array of capacity k
// lenght:   j - i
// capacity: k -i
// for slice

// slices and maps
slice := []int{10, 20, 30, 40}
// iterate over each element and display each value
for index, value := range slice {
	fmt.Printf("Index: %d Value: %d/n", index, value)
}
Output:
Index: 0 Value: 10
Index: 1 Value: 20
Index: 2 Value: 30
Index: 3 Value: 40

slice := []int{10, 20, 30, 40}
// iterate over each element and display each value and addresses
for index, value := range slice {
	fmt.Printf("Value: %d Value-Addr: %X ElemeAddr: %X\n", value, &value, &slice[index])
}
Output:
Value: 10 Value-Addr: 10500168 ElemAddr: 1052E100
Value: 20 Value-Addr: 10500168 ElemAddr: 1052E104
Value: 30 Value-Addr: 10500168 ElemAddr: 1052E108
Value: 40 Value-Addr: 10500168 ElemAddr: 1052E10C

slice := []int{10, 20, 30, 40}
// iterate over each element and display each value
for _, value := range slice {
	fmt.Printf("Value: %d\n", value)
}
Output:
Value: 10
Value: 20
Value: 30
Value: 40

slice := []int{10, 20, 30, 40}
// iterate over each element starting at el 3
for index := 2; index < len(slice); index++ {
	fmt.Printf("Index: %d Value: %d\n", index, slice[index])
}
Output:
Index: 2 Value: 30
Index: 3 Value: 40

// create a slice of a slice of int
slice := [][]int{{10}, {100, 200}}
// append the value of 20 to the 1st slice of int
slice[0] = append(slice[0], 20)

// allocate a slice of 1 mil int
slice := make([]int, 1e6)
// pass the slice to the func foo
func foo(slice)
// func foo accepts a slice of int and returns the slice back
func foo(slice []int) []int {
	...
	return slice
}
