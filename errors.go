package main

import (
	// "fmt"
	// "errors"
)

// 1
//Error Interface
// func sendSMSToCouple(msgToCustomer, msgToSpouse string) (float64, error) {
// 	costForCustomer, err := sendSms(msgToCustomer)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	costForSpouse, err := sendSms(msgToSpouse)
// 	if err != nil {
// 		return 0.0, err
// 	}
// 	return costForCustomer + costForSpouse, nil
// }
// func sendSms(message string) (float64, error) {
// 	const maxTextLen = 25
// 	const costPerChar = .0002
// 	if len(message) > maxTextLen {
// 		return 0.0, fmt.Errorf("Can't send texts over %v characters",maxTextLen)
// 	}
// 	return costPerChar * float64(len(message)), nil
// }
// func test(msgToCustomer, msgToSpouse string) {
// 	defer fmt.Println("==============")
// 	fmt.Println("Message for customer:",msgToCustomer)
// 	fmt.Println("Message for spouse:",msgToSpouse)
// 	totalCost, err := sendSMSToCouple(msgToCustomer, msgToSpouse)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Total cost:", totalCost)
// }

// 2
//Custom Error Interface Types
// type divideError struct {
// 	dividend float64
// }
// func (de divideError) Error() string {
// 	return fmt.Sprintf("Can not divide %v by zero", de.dividend)
// }
// func divide(dividend, divisor float64) (float64, error) {
// 	if divisor == 0 {
// 		return 0.0, divideError{dividend: dividend}
// 	}
// 	return dividend / divisor, nil
// }

//3
//Error Packages
// func divide(x, y float64) (float64, error) {
// 	if y == 0 {
// 		return 0.0, errors.New("Cannot divided by zero")
// 	}
// 	return x / y, nil
// }
// func test(x, y float64) {
// 	defer fmt.Println("==============")
// 	fmt.Printf("Dividing %.2f by %.2f\n",x,y)
// 	quotient, err := divide(x,y)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Qoutient: %.2f\n",quotient)
// }

//func main() {
	//Error Interface test function called
	//test("Hello, customer!", "Hello, spouse!")


	//Custom Error Interface Types test function called
	// fmt.Println(divide(8, 0))
	// fmt.Println(divide(8, 2))

	//Error Packages test function called
// 	test(8, 0)
// 	test(8, 2)
//}
