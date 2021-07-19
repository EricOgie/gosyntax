package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var age int = 32
var isOn bool = false

/*
* int int8 int16 int32 init64
* uint uint8 uint16 uint32 uint64 uintptr
* byte = uint8, rune = int32
* float32, float64
* complex64, complex128
 */

func main() {
	salary := 350000
	var name string = "Eric Ogie"
	fmt.Println("My name is "+name, age)
	fmt.Println("Salary is " + strconv.FormatInt(int64(salary), 10))
	fmt.Println(reflect.TypeOf(salary))
	fmt.Println(strconv.Itoa(salary))
	fmt.Printf("%T\n", age)
	fmt.Println(reflect.TypeOf(isOn))
}
