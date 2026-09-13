package main
import (
	"fmt"
	"errors"
)

// 1
//Arrays
func getMessageWithRetries() [3]string {
	return [3]string{
		"Click here to sign up",
		"Pretty please click here",
		"We beg you to sign up",
	}
}
func send(name string, doneAt int) {
	fmt.Printf("sending to %v ...",name)
	fmt.Println()
	messages := getMessageWithRetries()
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending "%v"`,msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("THey responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("Complete failure")
		}
	}
}

//Slices
const (
	planFree = "free"
	planPro = "pro"
)
func getMessageWithRetriesWithPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("Unsupported plan")
}
func getMessageWithRetries() [3]string {
	return [3]string {
		"Click here to sign up",
 		"Pretty please click here",
 		"We beg you to sign up",
	}
}
func test(name string, doneAt int, plan string) {
	defer fmt.Println("===================================================")
	fmt.Printf("sending to %v ...",name)
	fmt.Println()
 	messages, err := getMessageWithRetriesWithPlan(plan)
 	if err != nil {
 		fmt.Println("Error:", err)
 		return
 	}
 	for i := 0; i < len(messages); i++ {
 		msg := messages[i]
 		fmt.Printf(`sending "%v"`,msg)
 		fmt.Println()
 		if i == doneAt {
 			fmt.Println("They responded!")
 			break
 		}
 		if i == len(messages)-1 {
 			fmt.Println("No response")
		}
 	}
}

//Make keyword
func getMessageCost(messages []string) float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		messages := messages[i]
		cost := float64(len(messages)) * 0.01
		costs[i] = cost
	}
	return costs[len(costs)-1]
}

//Variadic functions and Spread Operator
func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}
func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n",len(nums))
	fmt.Printf("Bill for the month: %.2f\n",total)
	fmt.Println("======= END REPORT ========")
}

func main() {
	//Array send func called
	send("Bob",0)
	send("Alice",1)
	send("Sunny",2)
	send("Raj",3)

	//Slice test func called
	test("Sunny",3,planFree)
	test("Raj",3,planPro)
	test("Rani",3,"No plan")

//Make keyword test
messages := []string{
	"Click here to sign up",
	"Pretty please click here",
	"We beg you to sign up",
}
cost := getMessageCost(messages)
fmt.Printf("The cost of sending the messages is: $%.2f\n", cost)

Variadic functions and Spread Operator test func called
test(1.0,2.0,3.0)
test(1.0,2.0,3.0,4.0,5.0)
test(1.0,2.0,3.0,4.0,5.0,6.0,7.0,8.0,9.0,10.0)
}