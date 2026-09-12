package main

// import (
// 	"fmt"
// 	"sync"
// 	//"time"
// )

// func task(id int) {
// 	fmt.Println("Doing task", id);
// }
// func task(id int, w *sync.WaitGroup) {  //Always call WaitGroup as pointer, otherwise it will not work
// 	defer w.Done();  //.Done() is used to tell the WaitGroup that this goroutine is done executing 
// 	//defer is used to ensure that Done() is called when the function exits, even if it exits due to an error or panic
// 	fmt.Println("Doing task", id);
// }
// func main() {
	// for i := 0; i <= 10; i++ {
	// 	//go task(i);
	// 	go func(i int) { //Anonymous function is used to pass the value of i to the goroutine, otherwise it will always print 10
	// 		fmt.Println(i);
	// 	}(i);
	// }
	// time.Sleep(time.Second * 2);

// 	var wg sync.WaitGroup; //Initialize a WaitGroup to wait for all goroutines to finish before exiting the main function
// 	for i:= 0; i <= 10; i++ {
// 		wg.Add(1); //Add the number of goroutines to wait for, in this case 10
// 		go task(i, &wg);
// 	}
// 	wg.Wait(); //Wait for all goroutines to finish before exiting the main function
// }