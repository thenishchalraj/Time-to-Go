/*
The following rules apply to a switch statement −
    1. The expression used in a switch statement must have an variable of interface{} type.
    2. You can have any number of case statements within a switch. Each case is followed by the value to be compared to and a colon.
    3. The type for a case must be the same data type as the variable in the switch, and it must be a valid data type.
    4. When the variable being switched on is equal to a case, the statements following that case will execute. No break is needed in the case statement.
    5. A switch statement can have an optional default case, which must appear at the end of the switch. The default case can be used for performing a task when none of the cases is true. No break is needed in the default case.
*/

package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) { //type switch
	case int: //int type
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string: //string type
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default: //default
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(2)
	do("go")
	do(true)
}

/*
Output :

Twice 2 is 4
"go" is 2 bytes long
I don't know about type bool!

*/
