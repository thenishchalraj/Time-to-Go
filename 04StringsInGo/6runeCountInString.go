/*
The func RuneCountInString(s string) (n int) function of the utf8 package is used to find the length of the string.
This method takes a string as argument and returns the number of runes in it i.e. the length of the string.
This method also counts "space"
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func main() {

	word1 := "Learning Go"
	length(word1)
	word2 := "Counting characters in Strings"
	length(word2)
}

/*
Output :

length of Learning Go is 11
length of Counting characters in Strings is 30

*/
