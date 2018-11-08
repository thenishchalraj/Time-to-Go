/*
This is same as if statement in Go.
An if statement can be followed by an optional else statement, which executes when the boolean expression is false.
That is, if the condition in the if statement is false then else statement will be executed.
*/

package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 11

	/* check the boolean condition */
	if a == 10 {
		/* if condition is true then print the following */
		fmt.Printf("a is equal to 10\n")
	} else {
		/* if condition is false then print the following */
		fmt.Printf("a is not equal to 10\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

/*
Here in the above example if statement checks weather a is equal to 10 or not
In our case, a is equal to 11, so the else part will be executed
*/

/*
Output :

a is not equal to 10
value of a is : 100

*/
