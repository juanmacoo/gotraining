package main

import "fmt"

type Employee struct {
	firstName string
	lastName string
	id int
}

func main() {
	// Exercise 1
	greetings := []string{
		"Hello", "hola", "こんにちは", "안녕하세요", "Здравствуйте", "नमस्ते",
	}

	first := greetings[0:2]
	second := greetings[1:4]
	third := greetings[3:5]

	fmt.Println(greetings)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(third)

	// Exercise 2
	message := "Hi 👩 and 👨"

	mRune := []rune(message)

	fmt.Println(message)
	fmt.Println(string(mRune[3]))

	// Exercise 3
	employee1 := Employee{
		"Kt",
		"Joras",
		1,
	}

	employee2 := Employee{
		firstName: "Jin",
		lastName: "Yaloa",
		id: 2,
	}

	var employee3 Employee

	employee3.firstName = "Kate"
	employee3.lastName = "Tizor"
	employee3.id = 3

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
