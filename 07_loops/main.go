package main

import "fmt"

var colors = []string{"red", "blue", "green", "white", "nill"}

// ------ Simple for Loop ------------- //
func printColors() {
	for x := 0; x < len(colors); x++ {
		fmt.Println(colors[x])
	}
}

// --------- ForEach Loop ------------- //
func showColor() {
	for _, color := range colors {
		fmt.Println(color)
	}
}

// --------- FizzBuzz test ------------- //

func doFizzBuzz() {
	for x := 1; x < 100; x++ {
		if x%15 == 0 {
			fmt.Println("FIZZBUZZ")
		} else if x%5 == 0 {
			fmt.Println("BUZZ")
		} else if x%3 == 0 {
			fmt.Println("FIZZ")
		} else {
			fmt.Println(x)
		}
	}
}

func main() {
	// printColors()
	// showColor()
	doFizzBuzz()
}
