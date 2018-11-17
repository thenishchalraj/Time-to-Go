/*
if slatement is same as if used in english language i.e. a certain piece of code gets executed once a condition is
satisfied, otherwise not.
*/

package main

import "fmt"

func main() {

	//declaring local variable
	var a int = 5

	if a == 5 { //checking if a is equal to 5
		/* if condition is true then print the following */
		fmt.Printf("I have learned the if statement in Go\n")
	}
	fmt.Printf("value of a is : %d\n", a)
}

/*
Output :

I have learned the if statement in Go
value of a is : 5

*/
