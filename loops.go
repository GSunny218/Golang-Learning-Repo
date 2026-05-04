package main
//import "fmt"

// 1
// func bulkSend(numMessages int ) float64 {
// 	totalCost := 0.0
// 	for i := 0; i < numMessages; i++ {
// 		totalCost += 1.0 + (0.01 * float64(i))
// 	}
// 	return totalCost
// }
// func test(numMessages int) {
// 	fmt.Printf("Sending %v messages\n",numMessages)
// 	cost := bulkSend(numMessages)
// 	fmt.Printf("Bulk send complete cost = %.2f\n",cost)
// 	fmt.Println("========================")
// }


// 2
//Omitting conditions from a for loop
// func maxMessages(thresh float64) int {
// 	totalCost := 0.0
// 	for i := 0; ; i++ {
// 		totalCost += 1.0 + (0.01 * float64(i))
// 		if totalCost > thresh {
// 			return i
// 		}
// 	}
// }
// func test(thresh float64) {
// 	fmt.Printf("Thresh %.2f\n",thresh)
// 	max := maxMessages(thresh)
// 	fmt.Printf("Maximum messages that can be sent = %v\n",max)
// 	fmt.Println("=======================================")
// }

//For loop like while loop. Because Go don't have while loop, but for loop can be use like while loop
// func getMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
// 	actualCostInPennies := 1.0
// 	maxMessagesToSend := 0
// 	for actualCostInPennies <= float64(maxCostInPennies) {
// 		maxMessagesToSend++
// 		actualCostInPennies *= costMultiplier
// 	}
// 	return maxMessagesToSend
// }

//func main() {
	//Called test func of 1
	//test(50)

	//Called Omitting for loop test func
	//test(50)
//}