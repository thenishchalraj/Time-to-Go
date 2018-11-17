/*
Simple example of channel
*/
package main

import "fmt"

func main() {
	message := make(chan string) // making channel

	go func() { //function to assign value to channel "message"
		message <- "channel" //assigning value to channel "message"
	}()

	msg := <-message // value of message to msg
	fmt.Println(msg)
}

/*
Output :
channel
*/
