package main

import "fmt"

func main() {
	// Weird slice and subslice overwritting
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d") // x: ["a", "b", "c", "d"] cap:5
	y := x[:2] // y: ["a", "b"] cap: 5 Capacity is 3 because it grabs the capacity of the original (5) minus the initial offset of the subslice (0)
	z := x[2:] // z: ["c", "d"] cap: 3 Capacity is 3 because it grabs the capacity of the original (5) minus the initial offset of the subslice (2)
	fmt.Println(cap(x), cap(y), cap(z)) // x: 5, y: 5, z: 3
	y = append(y, "i", "j", "k") // x: ["a", "b", "i", "j", "k"] y: ["a", "b", "i", "j", "k"] z: ["i", "j"]
	x = append(x, "x") // x: ["a", "b", "i", "j", "k", "x"] y: ["a", "b", "i", "j", "x"] z: ["i", "j"]
	z = append(z, "z") // x: ["a", "b", "i", "j", "k", "z"] y: ["a", "b", "i", "j", "z"] z: ["i", "j", "z"]
	
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Using full slice expressions (x[0:2]) to avoid weird interactions
	xx := make([]string, 0, 5)
	xx = append(xx, "a", "b", "c", "d")
	yy := x[0:2:2]
	zz := x[2:4:4]
	fmt.Println(cap(xx), cap(yy), cap(zz))
	yy = append(yy, "i", "j", "k")
	xx = append(xx, "x")
	zz = append(zz, "z")
	
	fmt.Println("x:", xx)
	fmt.Println("y:", yy)
	fmt.Println("z:", zz)
}

