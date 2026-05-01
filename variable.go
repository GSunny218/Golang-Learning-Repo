package main

// import "fmt"

// func main() {
// 	/*Data Types
// 	int
// 	string
// 	bool
// 	uint32
// 	byte
// 	rune
// 	float64
// 	complex128
// 	*/

// 	/*First way to define variables*/
// 	var smsSendingLimit int
// 	var costPerSMS float64
// 	var hasPermission bool
// 	var username string
// 	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username);

// 	/*Second way to define variables*/
// 	penniesPerText := 2;            //Without defining with datatypes
// 	fmt.Printf("The type of penniesPerText is %T\n",penniesPerText)

// 	//Same Line Declaration
// 	averageOpenRate, displayMessege := .23,"is the average open rate";
// 	fmt.Println(averageOpenRate,displayMessege)

// 	//Type Sizes
// 	accountAge := 2.3;
// 	accountAgeInt := int(accountAge);
// 	fmt.Println("Your account has existed for",accountAgeInt,"years")

// 	//Constants
// 	const premiumPlanName = "Premium Plan"
// 	fmt.Println("Plan",premiumPlanName)

// 	//Formatting String
// 	const name = "Sunny"
// 	const openRate = 30.1
// 	msg := fmt.Sprintf("Hi %v your open rate is %.1f percent",name,openRate)
// 	fmt.Println(msg)

// 	//if statement with initial statement
// 	email := "gsunny@123"
// 	if length := len(email); length > 1 {
// 		fmt.Println("Email is invalid!")
// 	}
// }