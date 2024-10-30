package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)


// iota starts at zero like any other value
const (
	a = (2 * iota)
	b
	c
	d
)

const (
	c1 = iota
	c2
	c3 = iota
	c4
)

const (
	_ = iota
	KB float64 = 1 << (10 * iota)
	MB 
	GB 
	TB 
	PB 
	EB 
	ZB 
	YB 
)

// iota is used n constant blocks, or else it's just zero

func main() {
	// fmt.Println(x, y, z)
	// fmt.Println(a, b, c, d)
	// fmt.Println(c1, c2, c3, c4)

	var printFormat = "%f - %b\n"

	fmt.Printf(printFormat, KB,KB)
	fmt.Printf(printFormat, MB, MB)
	fmt.Printf(printFormat, GB, GB)
	fmt.Printf(printFormat, TB, TB)
	fmt.Printf(printFormat, PB, PB)
	fmt.Printf(printFormat, EB, EB)
	fmt.Printf(printFormat, ZB, ZB)
	fmt.Printf(printFormat, YB, YB)
}