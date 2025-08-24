package main

import "fmt"

const x = 10

func main() {
	var a int = x
	var b float64 = x
	var c byte = x

	fmt.Println(a,b,c)
	fmt.Printf("Type: %T\n", a)
	fmt.Printf("Type: %T\n", b)
	fmt.Printf("Type: %T\n", c)
}