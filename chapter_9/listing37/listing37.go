// sample program to show how different functions from the standart library use the io.Writer interface
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	// create a Buffer value and write a string to the buffer. Using the Write method that implements io.Writer
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// use Fprint to concatenate a string to the Buffer. Passing the address of a bytes.Buffer value for the io.Writer
	fmt.Fprintf(&b, "World!")

	// Write the content of a buffer to the stdout device Passing the address of a os.File value for io.Writer
	b.WriteTo(os.Stdout)
}
