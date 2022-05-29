## The bug

When sending data over a channel, the routine it's in waits until said data has been received. This means that the program pauses after `ch <- "Hello world!"` until something takes that data from the channel. This will never happen.

## Un-bugged Code

```go
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
```

## The fix
The program is supposed to print hello world; it now does that, since we spin up a second thread to send "Hello world!" into the channel. I assume that the usage of channels is essential, or the program would have been more simply written as "fmt.Println("Hello world!")".