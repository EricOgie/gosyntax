package main

import "fmt"

// ------- Long Declaration and assignment
// var emails = make(map[string]string)

func main() {

	// Short-hand declaration
	emails := make(map[string]string)

	emails["google"] = "eric@gmail.com"
	emails["yahoo"] = "ericogia@yahoo.com"
	emails["outlook"] = "eric-ogie.aghahowa@iubh.com"

	// Declaration and Assignment
	friends := map[int]string{0: "Chinebele Orji", 1: "Prince Aigbedion", 2: "Gule Phillip"}

	fmt.Println(emails["yahoo"])
	fmt.Println(emails["google"])
	fmt.Println(emails["outlook"])
	fmt.Println(friends[1])
	fmt.Println(friends[0])
	fmt.Println(friends[2])

}
