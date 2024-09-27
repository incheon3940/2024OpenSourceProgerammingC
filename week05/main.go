package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	var f float32 = 3.14

	///var i int = 33
	///i = 55

	i := 55
	fmt.Println(f, math.Ceil((3.14)))
	fmt.Println(strings.Title("Kim inha"))
	fmt.Printf("i는 %d\n", i)
	fmt.Print("i는 ", i, "\n")
	fmt.Println("i는", i)
}
