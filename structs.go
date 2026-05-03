package main

//import "fmt"

// import (
// 	"fmt"
// )

//messageToSend struct declared             //1
// type messageToSend struct {
// 	phonenumber int
// 	message     string
// }
// //test func declared
// func test(m messageToSend) {
// 	fmt.Printf("Sending message: '%s' to '%v'\n", m.message, m.phonenumber)
// 	fmt.Println("=============================")
// }

//Nested structs
// type messageToSend struct {
// 	message string
// 	sender user
// 	recipient user
// }
// type user struct {
// 	name string
// 	number int
// }
// func canSendMessage(mToSend messageToSend) bool {
// 	if mToSend.sender.name == "" {
// 		return false
// 	}
// 	if mToSend.recipient.name == "" {
// 		return false
// 	}
// 	if mToSend.sender.number == 0 {
// 		return false
// 	}
// 	if mToSend.recipient.number == 0 {
// 		return false
// 	}
// 	return true
// }

//Anonymous struct
// myCar := struct {
// 	Make string
// 	Model string
// } {
// 	Make: "tesla",
// 	Model: "model 3"
// }

//Nested Anonymous struct
// type car struct {
// 	Make string
// 	Model string
// 	Height int
// 	Width int
// 	Wheel struct {
// 		Radius int
// 		Material string
// 	}
// }

//Emmbedded structs
// type sender struct {
// 	rateLimit int
// 	user                   //This is an embedded struct, it allows us to access the fields of user struct directly from sender struct without having to specify the user struct name
// }
// type user struct {
// 	name string
// 	number int
// }
// func test(s sender) {
// 	fmt.Println("Sender name:",s.name)
// 	fmt.Println("Sender number:",s.number)
// 	fmt.Println("Sender rate limit:",s.rateLimit)
// 	fmt.Println("=============================")
// }

//Struct Methods
// type authenticationInfo struct {
// 	username string
// 	password string
// }
// func (authI authenticationInfo) getBasicAuth() string {  //This is a method that is associated with the authenticationInfo struct, it can be called on any variable of type authenticationInfo
// 	return fmt.Sprintf(
// 		"Authentication: Basic %s:%s",
// 		authI.username,
// 		authI.password,
// 	)
// }
// func test(authInfo authenticationInfo) {
// 	fmt.Println(authInfo.getBasicAuth())
// 	fmt.Println("=======================")
// }

//func main() {
	//test func with messageToSend called
	// test(messageToSend{
	// 	phonenumber: 148298442,
	// 	message:     "Thanks for joining us",
	// })
	// //test func with messageToSend called
	// test(messageToSend{
	// 	phonenumber: 148298442,
	// 	message:     "Love to have you aboard",
	// })
	// //test func with messageToSend called
	// test(messageToSend{
	// 	phonenumber: 148298442,
	// 	message:     "We're so excited to have you",
	// })

	//Emmbedded structs test func called
	// test(sender{
	// 	rateLimit: 100,
	// 	user: user{   //This is how we initialize the embedded struct, we have to specify the struct name and then the fields of that struct
	// 		name: "Sunny",
	// 		number: 148298442,
	// 	},
	// })

	//Struct Methods test func called
	// test(authenticationInfo{
	// 	username: "john_doe",
	// 	password: "secret123",
	// })
//}
