package main

import (
	"sync"
	"time"
)

func main() {

	c := make(chan int)
	var wg sync.WaitGroup
}

func doSomething(i int, wg *sync.WaitGroup) {
	time.Sleep(4 * time.Second)
}
