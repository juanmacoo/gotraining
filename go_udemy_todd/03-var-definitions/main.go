package main

import (
	"fmt"
)

var zeroValue int

func main() {
	shortDeclartion := "This is a short declaration"
	var a, b, c string = "a", "b", "c"
	var x, y, _ = 1, 2, 3

	fmt.Println(zeroValue)
	fmt.Println(shortDeclartion)
	fmt.Println(a, b, c)
	fmt.Println(x, y)
}
