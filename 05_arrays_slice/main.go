package main

import "fmt"

func main() {

	// Arrays
	var members [4]string
	// assignment
	members[0] = "EricOgie"
	members[1] = "keyleb01"
	members[2] = "Utibe"
	members[3] = "John"
	// Declaring Array and Assigning at same time
	watchers := [2]string{"Tochuwku", "Pepjoe"}

	fmt.Println(members)
	fmt.Println(watchers)

	fmt.Println(len(members)) // TO get the legnth
	fmt.Println(members[1:3]) // To get range of array items
}
