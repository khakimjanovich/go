// Here we declared a main package ( package is a way to group functions,
// and it's made up of all the files in the same directory).
package main

// Here we imported popular fmt package, which contains functions for formatting text,
// including printing to the console. This package is one of the standard library
// packages you got when you installed GO.
import "fmt"

// Here we implemented a main function to print a message to console. A main function
// executes by default when you run the main package.
func main() {
	fmt.Println("Hello, World!")
}
