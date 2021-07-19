package main

import "fmt"

// Pointers avail the opportunity to assign the memory location of a var

var a int = 6
var b = &a // the memory location of the var, a

// Re assigniing the pointer value using the position in memory

func main() {
	fmt.Println(b)
	fmt.Println(a)

	*b = 18
	fmt.Println(a)
}
