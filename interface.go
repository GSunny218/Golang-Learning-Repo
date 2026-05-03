package main

import (
	"fmt"
	//"time"
)

// 1
// func sendMessage(msg message) {  //This function takes an interface as a parameter, it can take any type that implements the message interface
// 	fmt.Println(msg.getMessage())
// }
// type message interface {   //This is an interface, it defines a set of methods that a type must implement to be considered as a message
// 	getMessage() string
// }
// type birthdayMessage struct {   //This struct implements the message interface by having a getMessage method
// 	birthdayTime time.Time
// 	recipientName string
// }
// func (bm birthdayMessage) getMessage() string {   //This is the getMessage method that implements the message interface for the birthdayMessage struct
// 	return  fmt.Sprintf("Hi %s, it's your birthday on %s",bm.recipientName,bm.birthdayTime)
// }
// type sendingReport struct {   //This struct also implements the message interface by having a getMessage method
// 	reportName string
// 	numberOfSends int
// }
// func (sr sendingReport) getMessage() string {       //This is the getMessage method that implements the message interface for the sendingReport struct
// 	return fmt.Sprintf("Report: %s, Number of sends: %d",sr.reportName,sr.numberOfSends)
// }
// func test(m message) {
// 	sendMessage(m)   //This function takes a message interface as a parameter, it can take any type that implements the message interface, and it calls the sendMessage function with that message
// 	fmt.Println("============================================")
// }

// 2
// type employee interface {
// 	getName() string
// 	getSalary() int
// }
// type contractor struct {
// 	name string
// 	hourlyPay int
// 	hoursPerYear int
// }
// func (c contractor) getName() string {
// 	return c.name
// }
// func (c contractor) getSalary() int {
// 	return c.hourlyPay * c.hoursPerYear
// }
// type fullTime struct {
// 	name string
// 	salary int
// }
// func (ft fullTime) getName() string {
// 	return  ft.name
// }
// func (ft fullTime) getSalary() int {
// 	return ft.salary
// }
// func test(e employee) {
// 	fmt.Println(e.getName(), e.getSalary())
// }

//3
//Multiple Interfaces
// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return 0.05 * float64(len(e.body))
// 	}
// 	return 0.01 * float64(len(e.body))
// }
// func (e email) print() {
// 	fmt.Println(e.body)
// }
// type expense interface {
// 	cost() float64
// }
// type printer interface {
// 	print()
// }
// type email struct {
// 	isSubscribed bool
// 	body string
// }
// func test(e expense, p printer) {
// 	fmt.Printf("Printing with cost: $%.2f ...\n",e.cost())
// 	p.print()
// }

//4
//Type Assertions
// func getExpenseReport(e expense) (string, float64) {
// 	em, ok := e.(email)
// 	if ok {
// 		return em.toAddress, em.cost()
// 	}
// 	s, ok := e.(sms)
// 	if ok {
// 		return s.toPhoneNumber, s.cost()
// 	}
// 	return "", 0.0
// }
// func (e email) cost() float64 {
// 	if !e.isSubscribed {
// 		return 0.05 * float64(len(e.body))
// 	}
// 	return 0.01 * float64(len(e.body))
// }
// func (i invalid) cost() float64 {
// 	return 0.0
// }
// type expense interface {
// 	cost() float64
// }
// type email struct {
// 	isSubscribed bool
// 	body         string
// 	toAddress    string
// }
// type sms struct {
// 	isSubscribed  bool
// 	body          string
// 	toPhoneNumber string
// }
// func (s sms) cost() float64 {
// 	if !s.isSubscribed {
// 		return float64(len(s.body)) * .1
// 	}
// 	return float64(len(s.body)) * .03
// }
// type invalid struct{}
// func test(e expense) {
// 	address, cost := getExpenseReport(e)
// 	switch e.(type) {
// 	case email:
// 		fmt.Printf("Sending email to %s with cost: $%.2f\n", address, cost)
// 		fmt.Println("===================================")
// 	case sms:
// 		fmt.Printf("Sending sms to %s with cost: $%.2f\n", address, cost)
// 		fmt.Println("===================================")
// 	default:
// 		fmt.Printf("Invalid expense type with cost: $%.2f\n", cost)
// 		fmt.Println("===================================")

// 	}
// }

// 5
//Type switch
// func getExpenseReport(e expense) (string, float64) {
// 	switch v := e.(type) {
// 	case email:
// 		return v.toAddress, v.cost()
// 	case sms:
// 		return v.toPhoneNumber, v.cost()
// 	default:
// 		return "", 0.0
// 	}
// }   This is how we can use a type switch to determine the type of the expense and get the appropriate report

func main() {
	// test(birthdayMessage{   //This is how we initialize the birthdayMessage struct, we have to specify the struct name and then the fields of that struct
	// 	birthdayTime: time.Date(2024, time.June, 15, 0, 0, 0, 0, time.UTC),
	// 	recipientName: "Sunny",
	// })
	// test(sendingReport{    //This is how we initialize the sendingReport struct, we have to specify the struct name and then the fields of that struct
	// 	reportName: "Monthly Report",
	// 	numberOfSends: 100,
	// })

	// test func called with contractor and fullTime structs
	// test(contractor{
	// 	name: "John Doe",
	// 	hourlyPay: 50,
	// 	hoursPerYear: 2000,
	// })
	// test(fullTime{
	// 	name: "Jane Smith",
	// 	salary: 50000,
	// })

	// test func called with email struct that implements both expense and printer interfaces
	// e := email{
	// 	isSubscribed: true,
	// 	body: "Hello there",
	// }
	// test(e,e)
	// e = email{
	// 	isSubscribed: false,
	// 	body: "I want my money back",
	// }
	// test(e,e)

	// test func called with email and sms structs that implement the expense interface, and an invalid struct that does not implement the expense interface
	// e := email{
	// 	isSubscribed: true,
	// 	body:         "Hello there",
	// 	toAddress:    "gsunny@123",
	// }
	// test(e)
	// e = email{
	// 	isSubscribed: false,
	// 	body:         "This meeting could have been an email",
	// 	toAddress:    "gsunny@123",
	// }
	// test(e)
	// s := sms{
	// 	isSubscribed:  true,
	// 	body:          "Hello there",
	// 	toPhoneNumber: "12432434",
	// }
	// test(s)
	// s = sms{
	// 	isSubscribed:  false,
	// 	body:          "This meeting could have been an email",
	// 	toPhoneNumber: "4313224134",
	// }
	// test(s)
	// test(invalid{})
}
