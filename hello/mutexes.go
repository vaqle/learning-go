package main

import (
	"fmt"
	"sync"
)

var (
	amount = 10
	mutex  sync.Mutex
)

//we use mutex to prevent the code from running concurrently meaning twice
//so lets say i had a function to add matches to the game list
//the function could be called from multiple goroutines so it would add the same match twice
//so we use a mutex to prevent that

func increase(val int, w *sync.WaitGroup) {
	mutex.Lock()   //we lock it from other goroutines to avoid  concurrent access
	amount += val  //increase the amount
	mutex.Unlock() //we unlock it to allow other goroutines to access it
	w.Done()       //we notify that we are done and we can move on
}

func withdraw(amount int, w *sync.WaitGroup) {
	mutex.Lock()     //we lock it from other goroutines to avoid  concurrent access
	amount -= amount //decrease the amount
	mutex.Unlock()   //we unlock it to allow other goroutines to access it
	w.Done()         //we notify that we are done and we can move on
}

func main() {

	fmt.Println("HELLO WORLD")
	amount := 400
	var wg sync.WaitGroup
	wg.Add(2)
	go increase(700, &wg)
	go withdraw(300, &wg)
	fmt.Println("amount:", amount)
}
