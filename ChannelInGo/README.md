Nowadays, we are using pipelines as ameans of transport i.e. you can send and recieve data from one part to other or simply you can send something from one place and receive it at some other place.

Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

For Example :
(The data flows in the direction of the arrow.)
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.