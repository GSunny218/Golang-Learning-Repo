package main

import (
	"fmt"
	"time"
	"math/rand"
)

func processNum(numChan chan int) { //For sending
	for  num := range numChan {
		fmt.Println("Processing number: ", num);
		time.Sleep(time.Second);
	}
	fmt.Println("Processing number: ", <- numChan);
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2;
	result <- numResult;
}

//Goroutine synchronizer
func task(done chan bool) {
	defer func() {done <- true}();
	fmt.Println("Pocessing...");
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() {done <- true}();
	for email := range emailChan {
		fmt.Println("Sending email:", email);
		time.Sleep(time.Second);
	}
}

func main() {
	messageChan := make(chan string);  //Create a channel of type string
	messageChan <- "ping";  // Send a message to the channel
	mainMessage := <-messageChan;  // Receive a message from the channel
	fmt.Println(mainMessage);  // Print the message received from the channel

	numChan := make(chan int);
	go processNum(numChan);
	for {   //Infinite loop
		numChan <- rand.Intn(100);  //generate random numbers between 0 to 100
	}
	//numChan <- 5;
	time.Sleep(time.Second * 2);

	result := make(chan int);
	go sum(result, 4, 5);
	res := <- result;
	fmt.Println(res);

	done := make(chan bool);
	go task(done);
	<- done; //block

	emailChan := make(chan string, 100)
	done := make(chan bool);
	go emailSender(emailChan, done);
	for i := 1; i < 100; i++ {
		emailChan <- fmt.Sprintf("%d@gamil.com", i);
	}
	emailChan <- "example@123.com";
	emailChan <- "test@123.com";
	fmt.Println("Done Sending...");
	close(emailChan); //Closing the channels. It's very important to close the channels after sending all the data. Otherwise, the receiver will be blocked forever.
	<- done; //block
	fmt.Println(<-emailChan);
	fmt.Println(<-emailChan);

	chan1 := make(chan int)
	chan2 := make(chan string)
	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "Hello"
	}()
	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2", chan2Val)
		}
	}
}
