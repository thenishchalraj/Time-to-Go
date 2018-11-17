/*
The following rules apply to a select statement −
    1. You can have any number of case statements within a select. Each case is followed by the value to be compared to and a colon.
    2. The type for a case must be the a communication channel operation.
    3. When the channel operation occured the statements following that case will execute. No break is needed in the case statement.
    4. A select statement can have an optional default case, which must appear at the end of the select. The default case can be used for performing a task when none of the cases is true. No break is needed in the default case.
*/
package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second) //Sleep function called
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second) //Sleep function called
		c2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*
Output :

received one
received two

*/
