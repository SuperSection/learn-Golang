package main

// int, float64, bool, string, rune
// Default type 0, 0.0, false, ""

import (
	"fmt"
	"reflect"
)

// jwtToken  := 30000 -> this short variable declaration is not allowed outside of function or globally

const LoginToken string = "paglirpagol"	// public

func main() {

	fmt.Println(reflect.TypeOf(24))
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("Love You"))
	fmt.Println(reflect.TypeOf(`❣️`))
	
	var username string = "pagol"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var smallVal uint8 = 255
	fmt.Println(smallVal) 
	fmt.Printf("Variable is of type: %T \n", smallVal)
	
	var smallFloat32 float32 = 255.44291869735
	var smallFloat64 float64 = 255.44291869735
	fmt.Println(smallFloat32)
	fmt.Println(smallFloat64)
	fmt.Printf("Variable is of type: %T \n", smallFloat32)
	fmt.Printf("Variable is of type: %T \n", smallFloat64)

// default value and some aliases
	var anotherInteger int
	fmt.Println(anotherInteger)
	fmt.Printf("Variable is of type: %T \n", anotherInteger)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)
	
// implicit type
	var website = "supersection.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)
	// website = 5 -> can't change it to a differen type when the variable is already treated as one type (after initialization)

// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
