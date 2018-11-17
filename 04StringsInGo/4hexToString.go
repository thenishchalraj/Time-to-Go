/*
We have converted the given string into hexa decimal in the previous example. But can we do the reverse process ?
The answer is yes, We can convert the hexadecimal to string using byte.
Following is the example of the same :
*/

package main

import (
	"fmt"
)

func main() {
	byteToString := []byte{0x49, 0x20, 0x61, 0x6d, 0x20, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x20, 0x47, 0x6f} //byte is used assign the hexadecimal value of string
	str := string(byteToString)
	fmt.Println(str)
}

/*
Output :

I am learning Go

*/
/*
To Do :
Try to put decimal vaule in place of hexadecimal and find the output.
(You can get the decimal and hexadecimal value from previous program)
*/
