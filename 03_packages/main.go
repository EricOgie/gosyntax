package main

import (
	"fmt"
	"math"

	"github.com/EricOgie/gosyntax/03_packages/utils"
)

func main() {
	fmt.Println(math.Floor(3.6), math.Ceil(3.6))
	fmt.Println(utils.RunBack("hello"))

}
