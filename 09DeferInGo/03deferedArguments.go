/*
The arguments of a deferred function are evaluated when the defer statement is executed and not
when the actual function call is done. In simple way, if a defered function is having an argument
then the value of the argument will be same as that value at that time and not the value just before
the return statement. Let's understand this by an example :
*/
package main

import (
	"fmt"
)

func deferCall(a int) {
	fmt.Println("value of a in deferred function", a)
}
func main() {
	a := 5
	defer deferCall(a) //here value of a is 5 so 5 will be passed whenever the deferCall() will be called
	a = 10
	fmt.Println("value of a before deferred function call", a)
}

/*
Output :

value of a before deferred function call 10
value of a in deferred function 5
*/
