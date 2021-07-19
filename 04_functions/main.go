package main

import "fmt"

func greet(name string) string {
	return "Welcome " + name + ", it's good to have you here"
}

func main() {
	fmt.Println(greet("Eric Ogie"))
}
