/*
In last example, We were trying to print characters assuming that each code point will be one byte long which is wrong.
In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this.
This is where rune saves us.
*/
/*
A rune is a builtin type in Go and it's the alias of int32. rune represents a Unicode code point in Go.
It does not matter how many bytes the code point occupies, it can be represented by a rune.
Lets modify the previous program to print characters using a rune.
*/

package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func main() {
	name := "I am learning Go"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n\n")
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}

/*
Output :

49 20 61 6d 20 6c 65 61 72 6e 69 6e 67 20 47 6f
I   a m   l e a r n i n g   G o

53 65 c3 b1 6f 72
S e ñ o r
*/
