/*
Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.

v, ok := <- ch

In the above statement ok is true if the value was received by a successful send operation to a channel.
If ok is false it means that we are reading from a closed channel. The value read from a closed channel will be the zero value of the channel's type.
For example if the channel is an int channel, then the value received from a closed channel will be 0.
*/

package main

import (
	"fmt"
)

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}
func main() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}
}

/*
In the program above, the producer Goroutine writes 0 to 9 to the chnl channel and then closes the channel.
The main function has an infinite for loop in line no.16 which checks whether the channel is closed using the variable ok in line no. 18.
If ok is false it means that the channel is closed and hence the loop is broken. Else the received value and the value of ok is printed.
This program prints :

Received  0 true
Received  1 true
Received  2 true
Received  3 true
Received  4 true
Received  5 true
Received  6 true
Received  7 true
Received  8 true
Received  9 true
*/
