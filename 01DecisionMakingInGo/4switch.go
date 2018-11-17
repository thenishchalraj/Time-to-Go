/*
Switch case statements are a substitute for long if statements that compare a variable to several integral values
    1. The switch statement is a multiway branch statement. It provides an easy way to dispatch execution to different parts of code based on the value of the expression.
    2. Switch is a control statement that allows a value to change control of execution.
*/
/*
In Go programming, switch statements are of two types −
    1. Expression Switch − In expression switch, a case contains expressions, which is compared against the value of the switch expression.
    2. Type Switch − In type switch, a case contain type which is compared against the type of a specially annotated switch expression.
*/
/*
Expression Switch :

The following rules apply to a switch statement −
    1. The expression used in a switch statement must have an integral or boolean expression, or be of a class type in which the class has a single conversion function to an integral or boolean value. If the expression is not passed then the default value is true.
    2. You can have any number of case statements within a switch. Each case is followed by the value to be compared to and a colon.
    3. The constant-expression for a case must be the same data type as the variable in the switch, and it must be a constant or a literal.
    4. When the variable being switched on is equal to a case, the statements following that case will execute. No break is needed in the case statement.
    5. A switch statement can have an optional default case, which must appear at the end of the switch. The default case can be used for performing a task when none of the cases is true. No break is needed in the default case.
*/

package main

import "fmt"

func main() {
	/* local variable definition */
	var grade string = "B"
	var marks int = 90

	switch marks { //switch on the basus of marks
	case 90:
		grade = "A" //assigning grade
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	switch { //switch on the basis of grade
	case grade == "A":
		fmt.Printf("Excellent!\n")
	case grade == "B", grade == "C":
		fmt.Printf("Well done\n")
	case grade == "D":
		fmt.Printf("You passed\n")
	case grade == "F":
		fmt.Printf("Better try again\n")
	default:
		fmt.Printf("Invalid grade\n")
	}
	fmt.Printf("Your grade is  %s\n", grade)
}

/*
Output :

Excellent!
Your grade is  A

*/

/*
To Do :
This switch case is only valid for marks as a multiple of 10. Try to make it for all values i.e. including non multiples of 10
*/
