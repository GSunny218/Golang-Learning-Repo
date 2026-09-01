package main;

// import (
// 	"fmt"
//  	"sync"
// );

// type post struct {
// 	views int;
// 	mu sync.Mutex;
// }

// func (p *post) inc(wg *sync.WaitGroup) {
// 	defer func() {
// 		wg.Done();
// 		p.mu.Unlock(); //Unlock the mutex to allow other goroutines to access the critical section
// 	}();
// 	p.mu.Lock(); //Lock the mutex to ensure that only one goroutine can access the critical section at a time
// 	p.views++; //Increment the views count
// }

// func main() {
// 	var wg sync.WaitGroup;
// 	myPost := post{views: 0};
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1);
// 		go myPost.inc(&wg);
// 	}
// 	wg.Wait();
// 	// myPost.inc();
// 	// myPost.inc();
// 	fmt.Println(myPost.views);
// }