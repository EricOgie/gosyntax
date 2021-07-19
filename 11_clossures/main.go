package main

import "fmt"

/**
*  Clossures are functions that take other functions as inpute or return them
*  They particularly known to keep a reference to a variable outside of their scope
*  They are very handy for algorithm with level addition or subtraction
**/

//  Defining function with level addition

func addFunc() func(int) int {
	sum := 1
	return func(x int) int {
		sum += x
		return sum
	}

}

//  Defining function with level Subtraction

func subFunc() func(int) int {
	bigNumber := 30
	return func(i int) int {
		bigNumber -= i
		return bigNumber
	}
}

// Run amin. go

func main() {
	seriesVal := addFunc()
	for x := 0; x < 10; x++ {
		fmt.Println(seriesVal(x))
	}

	dec := subFunc()
	for x := 0; x < 10; x++ {
		fmt.Println(dec(x))
	}
}
