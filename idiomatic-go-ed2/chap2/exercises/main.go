package main

import (
	"fmt"
	"math"
)

func main() {
	// Exercise 1
	i := 20
	var f = float64(i)

	fmt.Println(i)
	fmt.Println(f)

	// Exercise 2
	const value = 20
	var ii int = value
	var ff float64 = value

	fmt.Println(ii)
	fmt.Println(ff)

	//Exercise 3
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	fmt.Println(b + 1)
	fmt.Println(smallI + 1)
	fmt.Println(bigI + 1)
}