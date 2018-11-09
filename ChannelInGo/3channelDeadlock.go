/*
Deadlock :
One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.
*/

package main

func main() {
	ch := make(chan int)
	ch <- 5
}

/*
In the program above, a channel ch is created and we send 5 to the channel in line ch <- 5.
In this program no other Goroutine is receiving data from the channel ch.
Hence this program will panic with the following runtime error :

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	/home/sumit/Time-to-Go/ChannelInGo/3channelDeadlock.go:11 +0x50
exit status 2
*/
