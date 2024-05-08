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
  score := 0
  num_Questions := 2

	fmt.Printf("What is better the RTX 3080 or the RTX 3090? ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)
	fmt.Printf("You answered: %v %v \n", answer, answer2)

	if answer + " " + answer2 == "RTX 3090" {
		fmt.Println("Correct!")
		score ++
	}  else if answer + " " + answer2 == "rtx 3090" {
		fmt.Println("Correct!")
		score ++
	}else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many cores does the Ryzen 9 3900x have? ")
	var coresAnswer int
	fmt.Scan(&coresAnswer)

	if coresAnswer == 12 {
		fmt.Println("Correct!")
		score ++
	} else {
		fmt.Println("Incorrect!")
	}
	fmt.Printf("You got a score of %v out of %v \n", score, num_Questions)
	percent := (score / num_Questions) * 100
	fmt.Printf("You got a score of %v%% \n", percent)
}
