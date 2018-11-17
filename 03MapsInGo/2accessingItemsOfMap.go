/*
So, till now we have added the keys and values in the map.
Lets learn how to retrieve them.
'map[key]' is the syntax to retrieve elements of a map.

What will happen if a element is not present?
The map will return the zero value of the type of that element.

What if we want to know whether a key is present in a map or not?

"value, ok := map[key]"

The above is the syntax to find out whether a particular key is present in a map.
If ok is true, then the key is present and its value is present in the variable value, else the key is absent.
*/

package main

import (
	"fmt"
)

func main() {
	studentRoll := map[int]string{
		1: "Sumit",
		2: "Amit",
	}
	studentRoll[3] = "Rahul"
	roll := 1
	fmt.Println("Name of roll number ", roll, "is", studentRoll[roll])

	//trying to access value of key that doesn't exist
	fmt.Println("Name of roll 4 is", studentRoll[4])

	//finding keys in map
	newRoll := 3
	value, ok := studentRoll[newRoll] // ok will be true if newRoll exists otherwise false
	if ok == true {
		fmt.Println("Name of roll number ", newRoll, "is", value)
	} else {
		fmt.Println(newRoll, "not found")
	}
}

/*
Output :

Name of roll number  1 is Sumit
Name of roll 4 is
Name of roll number  3 is Rahul

*/
