package main

import (
	"fmt"
	"reflect"
)

// var name string = "Eric"
var name int = 44

// ----------  Switch control --------------- //

var colors = []string{"red", "blue", "green", "white", "nill"}

func runSwitch(color string) string {
	switch color {
	case "red":
		return "RED"
	case "blue":
		return "BLUE"
	case "green":
		return "GREEN"
	case "white":
		return "WHITE"
	default:
		return "No Color"
	}
}

func main() {

	if reflect.TypeOf(name).Kind() == reflect.String {
		fmt.Println("It is String, boss")
	} else if reflect.TypeOf(name).Kind() == reflect.Int {
		fmt.Println("It is an Integer, Baba")
	}

	fmt.Println("The color chosen is " + runSwitch(colors[3]))
}
