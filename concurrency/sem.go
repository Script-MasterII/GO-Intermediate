package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething_(i, &wg, c)
	}
	wg.Wait()
}

func doSomething_(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("id: %d started\n", i)
	time.Sleep(4 * time.Second)
	fmt.Printf("id: %d finished\n", i)
	<-c
}
