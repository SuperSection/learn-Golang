package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	
	welcome := "Welcome to user-input"
	fmt.Println(welcome)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || comma err
	input, _ := reader1.ReadString('\n')	// _ -> blank identifier
	fmt.Println("Thanks for rating,", input)

	fmt.Printf("Type of this rating is %T\n", input)


// another example
	fmt.Println("What's your name?")
	reader2 := bufio.NewReader(os.Stdin)
	name, err := reader2.ReadString('\n')	// best practice
	if err == nil {
		fmt.Println("Hello", name)
	} else {
		log.Fatal(err)
	}

}