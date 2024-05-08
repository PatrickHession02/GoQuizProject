package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz!")
	fmt.Println("What is your name?")
	var name string 
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, let's get started! \n", name)

	fmt.Printf("Enter your age ")
	var age uint
	fmt.Scan(&age)


	if age >= 18 {
		fmt.Println("You are old enough to take the quiz!")
	} else {
		fmt.Println("You are not old enough to take the quiz!")
		return
	}

	fmt.Printf("What is better the RTX 3080 or the RTX 3090? ")

	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Printf("You answered: %v %v \n", answer, answer2)
}
