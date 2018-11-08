/*
There are certain cases when you you want a particular statement to be executed on if certain condition is satisfied.
To do this, you can use one if or else if statement inside another if or else if statement(s).

For example : In order to vote for India's PM, a person must be a citizen of India and then he/she must be above 18
ears of age. So, firstly you have to check if he is acitizen of india, if no then he has no rights to vote, if yes
then check if he is above 18 then he can vote.
*/

package main

import "fmt"

func main() {
	/* local variable definition */
	var citizen string = "indian"
	var age int = 20

	/* check the boolean condition */
	if citizen == "indian" {
		/* if condition is true then check the following */
		if age >= 18 {
			/* if condition is true then print the following */
			fmt.Printf("You are eligible to vote\n")
		} else {
			fmt.Printf("You are not eligible to vote as your age is less than 18")
		}
	} else {
		fmt.Printf("You are not eligible to vote as you are not a citizen of India")
	}
}

/*
Output :

You are eligible to vote

*/
