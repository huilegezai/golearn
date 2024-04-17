// channel
package main

import "fmt"

func main() {

	message := make(chan string) // Create a new channel

	go func() {
		message <- "ping" // send value into channel
	}()

	msg := <-message // receives a value
	fmt.Println(msg)

	// channel buffering
	me := make(chan string, 2)
	me <- "ss"
	me <- "dd"
	fmt.Println(<-me)
	fmt.Println(<-me)
}
