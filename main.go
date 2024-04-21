package main

import (
	"fmt"
	"sync"
)

func printMsg(wg *sync.WaitGroup, msg string) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	var wg sync.WaitGroup

	//add the number of go routines to wait for
	wg.Add(3)

	go printMsg(&wg, "Hello")
	go printMsg(&wg, "Go")
	go printMsg(&wg, "Waitgroup")

	wg.Wait() //wait for all the go routines to finish
	fmt.Println("all messages are printed")
}
