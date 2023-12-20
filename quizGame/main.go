package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %v, welcome to the game!\n", name)

	fmt.Print("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Ya, you can play!")
	} else {
		fmt.Println("You can't play!")
		return
	}

	score := 0
	numOfQuestions := 2

	fmt.Print("Which one is better, the RTX 3080 or RTX 3090? ")
	var answer_part1 string
	var answer_part2 string
	fmt.Scan(&answer_part1, &answer_part2)

	if answer_part1+" "+answer_part2 == "RTX 3090" {
		fmt.Println("Correct!")
		score++
	} else if answer_part1+" "+answer_part2 == "rtx 3090" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Print("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Printf("\nYou scored %v out of %v.", score, numOfQuestions)
	percentScore := (float64(score) / float64(numOfQuestions)) * 100
	fmt.Printf("\nYou scored: %v%%.", percentScore)
}
