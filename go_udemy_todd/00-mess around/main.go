package main

import (
	"fmt"
    "strconv"
)

func main() {
    var name, age = "Kim", 22

    fmt.Printf("Hello %v, you're %v in this World! 🥳", name, age)
    fmt.Println(`
This is some text
And this is another text
A couple of lines of text
    `)

    // By default if you declare a variable without assigning a value,
    // it will be assigned to the zero value of the type. For int it is 0
    var g int
    fmt.Println(g)

    // By default if you declare a variable without assigning a value,
    // it will be assigned to the zero value of the type. For bool it is false
    var someBool bool
    fmt.Println(someBool)

    // By default if you declare a variable without assigning a value,
    // it will be assigned to the zero value of the type. For string it is ""
    var sometString string
    fmt.Println(sometString)

    // Short variable declaration
    shortVar := "This is a short variable declaration"
    fmt.Println(shortVar)

    a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
    fmt.Printf("Numbers in binary: %b, %b, %b, %b, %b, %b\n", a, b, c, d, e, f)
    fmt.Printf("Numbers in hexadecimal: %#x, %#x, %#x, %#x, %#x, %#x\n", a, b, c, d, e, f)

    // Conversions
    var i string = "42"
    j, _ := strconv.ParseInt(i, 10, 64)
    fmt.Printf("Converted number is: %v. The type is %T", j, j)
}