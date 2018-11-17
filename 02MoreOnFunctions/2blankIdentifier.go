/*
Go has built-in support for multiple return values. Sounds interesting ? because almost every langauge can return single
value. But Go has multiple return values to identify the error and output at the same time. For example, your program is
correct upto some place and after that there is some error, then at these type of situations multiple return values come
into existance.
*/

/*
Blank Indetifier :
If you only want a subset of the returned values, use the blank identifier _. i.e. if you want to get only some values
out of many returned values the use blank identifier.
*/

package main

import "fmt"

func multipleReturn() (string, int) { // function that takes no argument but returns multiple values of type 'string' and 'int'
	return "Go", 9
}
func main() {
	a, b := multipleReturn() // 'a' will store string value and 'b' will store int value
	fmt.Println(a)           // print value of 'a'
	fmt.Println(b)           // print value of 'b'

	_, c := multipleReturn() // use of blank identifier. Here we are getting 'c' only
	fmt.Println(c)
}

/*
Output :
Go
9
9
*/
