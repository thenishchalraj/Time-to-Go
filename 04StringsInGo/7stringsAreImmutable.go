/*
Immutable means you can't change it.
Strings are immutable in Go. Once a string is created it's not possible to change it.
*/
package main

import (
	"fmt"
)

func immutableToMutate(s string) string {
	s[0] = 'a' //trying to change the character at 0 index
	return s
}
func main() {
	str := "go"
	fmt.Println(immutableToMutate(str))
}

/*
Output :

# command-line-arguments
./7stringsAreImmutable.go:12:7: cannot assign to s[0]

*/
/*
Note:
You can change the string by using slice of string. So try to make it mutable.
We will learn this in the next part.
*/
