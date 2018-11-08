/*
Whenever we send something through channel then there must be areceiver to recieve it. This situation is called Unbuffered case.
By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding
receive (<- chan) ready to receive the sent value. Buffered channels accept a limited number of values without
a corresponding receiver for those values.
*/

package main

import "fmt"

func main() {
	messages := make(chan string, 3) //Here we make a channel of strings buffering up to 3 values.

	messages <- "buffered"
	messages <- "channel"
	messages <- "example"

	fmt.Println(<-messages) //values are received in the order they are assigned
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

/*
Output :
buffered
channel
example
*/
