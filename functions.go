package main

import (
	"errors"
	"fmt"
)

// This are function signatures
//
//	|           |
func concat(s1 string, s2 string) string {
	return s1 + s2
}
func main() {
	//concat function called
	//fmt.Println(concat("Sunny ","Hi"))

	//incrementSends function called
	sendsSoFar := 430
	const sendsToAdd = 25
	//sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messeges!")

	//getName function called
	//This _ ignores the second returned values from function
	//firstName, _ := getNames()
	//fmt.Println("Welcome",firstName)

	//test function called
	// test(4)
	// test(10)
	// test(22)

	//calculator function called
	//fmt.Println(calculator(2, 0))
	//fmt.Println(calculator(4, 2))
}

// Passing variables by Values
func incrementSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}

// Ignoring return values
func getNames() (string, string) {
	return "Sunny", "Gupta"
}

// Named Return Values
func yearUntilEvents(age int) (yearsUntilAdult int, yearsUntilDrinking int, yearsUntilCarRents int) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRents = 25 - age
	if yearsUntilCarRents < 0 {
		yearsUntilCarRents = 0
	}
	return
}
func test(age int) {
	fmt.Println("Age:", age)
	yearUntilAdult, yearUntilDrinking, yearsUntilCarRents := yearUntilEvents(age)
	fmt.Println("You're an adult in", yearUntilAdult, "years")
	fmt.Println("You can drink in", yearUntilDrinking, "years")
	fmt.Println("You can rent a car in", yearsUntilCarRents, "years")
	fmt.Println("=================================================")
}

func calculator(a, b int) (mul, div int, err error) {
	if b == 0 {
		return 0, 0, errors.New("Can't divide by zero!")
	}
	mul = a * b
	div = a / b
	return mul, div, nil
}