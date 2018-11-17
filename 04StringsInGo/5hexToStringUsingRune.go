/*
Can we convert the hexadecimal to string using 'rune' (learn more about rune in '3rune.go')
Let's try for the same :
*/

package main

import (
	"fmt"
)

func main() {
	runeSlice := []rune{0x49, 0x20, 0x61, 0x6d, 0x20, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x47, 0x6f}
	str := string(runeSlice)
	fmt.Println(str)
}

/*
Output :

I am learning Go

*/
