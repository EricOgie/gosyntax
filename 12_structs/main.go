package main

import (
	"fmt"
	"strconv"
)

// Person Struc
type Person struct {
	firstName string
	lastName  string
	school    string
	country   string
	age       int
}

// Language Struc
type Language struct {
	name       string
	category   string
	appication []string
}

type rect struct {
	legth int
	width int
}

type circle struct {
	radius int
}

// Value reciever
func (p Person) holla() string {
	return "How far, baba? My name na " + p.firstName + " and I don pluck " + strconv.Itoa(p.age) + " years"
}

// Pointer Receiver. This is use for attribute manaipulation

func (p *Person) doBday() {
	p.age++
}

func main() {
	appArry := []string{"Android", "web"}
	eric := Person{firstName: "Eric", lastName: "Ogie", school: "UNIBEN", country: "Nigeria", age: 32}
	lan := Language{"Java", "Programming", appArry}
	fmt.Println(lan)
	eric.doBday() // will add a year to Eric's aga.
	fmt.Println("Hi, I am " + eric.firstName + " " + eric.lastName + " and I am a graduate of " + eric.school)
	fmt.Println(eric.holla())
}
