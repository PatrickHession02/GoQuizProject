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
	fmt.Println(age >= 18)
}
