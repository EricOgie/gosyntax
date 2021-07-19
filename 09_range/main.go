package main

import (
	"fmt"
	"strconv"
)

/*
* We use range for looping
**/

func printD(data map[int]string) int {
	cum := 0
	for x, friend := range data {
		cum += x
		fmt.Println(friend)
	}
	return cum

}

func main() {

	emails := make(map[string]string)

	emails["google"] = "eric@gmail.com"
	emails["yahoo"] = "ericogia@yahoo.com"
	emails["outlook"] = "eric-ogie.aghahowa@iubh.com"

	friends := map[int]string{0: "Chinebele Orji", 1: "Prince Aigbedion", 2: "Gule Phillip"}

	total := printD(friends)
	fmt.Println("Total of Friends index is = " + strconv.Itoa(total))

	for x, email := range emails {
		fmt.Printf("%s - Friend: %s\n ", x, email)
	}

}
