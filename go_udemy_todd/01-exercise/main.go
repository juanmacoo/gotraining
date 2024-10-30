package main

import "fmt"

func cool(a, b int) (x, y int){
    x = a + y
    y = a - b
    return
}

func main(){
    fmt.Println(cool(1, 2))
}