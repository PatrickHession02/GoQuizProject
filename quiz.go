package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz!")
	fmt.Println("What is your name?")
	var name string 
	fmt.Scan(&name)
	fmt.Printf("Hello, %v, let's get started!", name)
}
