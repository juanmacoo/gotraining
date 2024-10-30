package main

import "fmt"

var numbers []int = make([]int, 10)

func main() {
	for index := range numbers {
		if index % 2 == 0 {
			fmt.Println(index, "is even")
		} else {
			fmt.Println(index, "is odd")
		}
	}
}