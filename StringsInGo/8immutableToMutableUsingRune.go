/*
Inorder to change the string i.e. immutable to immutable, you have to make the slice of it using runes
Strings are converted to a slice of runes.
Then that slice is mutated with whatever changes needed and converted back to a new string.
*/

package main

import (
	"fmt"
)

func immutableToMutable(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	str := "go"
	fmt.Println(immutableToMutable([]rune(str)))
}

/*
Output :

ao

*/
