package main

import (
	"fmt"
	"math"
)

func main(){
	aString := "String"
	var number int = 64
	var float float64 = 32.56

	stringFormat := "%v is of type %T\n"

	fmt.Printf(stringFormat, aString, aString)
	fmt.Printf(stringFormat, number, number)
	fmt.Printf(stringFormat, float, float)

	decimal := 747
	binary := 911
	hexadecimal := 90210

	fmt.Printf("%d is of type  %T\n", decimal, decimal)
	fmt.Printf("%b is of type  %T\n", binary, binary)
	fmt.Printf("%#x is of type  %T\n", hexadecimal, hexadecimal)

	largestInt := math.MaxInt8
	largestUnsigned := math.MaxUint8

	fmt.Printf("The largest int8 is: %v\n", largestInt)
	fmt.Printf("The largest uint8 is: %v", largestUnsigned)

}