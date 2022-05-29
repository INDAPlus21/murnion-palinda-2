package main

import "fmt"

func main() {
	channel := make(chan string)
	var sendString = func(channel chan<- string, text string) {
		channel <- text
	}
	go sendString(channel, "Hello world!")
	fmt.Println(<-channel)
}
