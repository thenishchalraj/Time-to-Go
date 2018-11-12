/*
Passing pointer to function.
*/

package main

import (
	"fmt"
)

func pointerToFunction(val *int) { //passing pointer to function
	*val = 10 //changing value of a
}
func main() {
	a := 20
	fmt.Println("value of a before function call is", a)
	b := &a
	pointerToFunction(b)
	fmt.Println("value of a after function call is", a)
}

/*
Output :

value of a before function call is 20
value of a after function call is 10

*/
