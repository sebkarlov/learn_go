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

// 