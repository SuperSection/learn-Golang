package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.October, 5, 5, 30, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("Monday 15:04:05 01-02-2006"))
}

// Reference link: https://pkg.go.dev/time
