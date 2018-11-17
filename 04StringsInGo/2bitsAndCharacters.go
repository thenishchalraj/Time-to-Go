/*
In this program we will try to access each bit of the string and print the respective charaters
*/
package main

import (
	"fmt"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ { //len() is a function used to find the length of a string
		fmt.Printf("%x ", s[i]) //printing hexadecimal
	}
}

func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "I am learning Go"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}

/*
Output :

49 20 61 6d 20 6c 65 61 72 6e 69 6e 67 20 47 6f
I   a m   l e a r n i n g   G o

*/

/*
Try to run the same code by using string as : "Señor"
NOTE : We are trying to print characters assuming that each code point will be one byte long which is wrong.
In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this.
This is where rune saves us.
*/
