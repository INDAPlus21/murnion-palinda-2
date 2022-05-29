package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	wg.Add(1)
	go Print(ch)
	for i := 1; i <= 11; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}

func Print(channel <-chan int) {
	defer wg.Done()
	for n := range channel {
		time.Sleep(10 * time.Millisecond)
		fmt.Println(n)
	}
}
