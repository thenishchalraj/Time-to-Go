/*
studentRoll := make(map[int]string)
The above line of code creates a map named studentRoll which has int keys and string values.

Note : The zero value of a map is nil. If you try to add items to nil map, a run time panic will occur.
Hence the map has to be initialized using make function.
*/

package main

import (
	"fmt"
)

func main() {
	studentRoll := make(map[int]string)                   //creating map named studentRoll
	studentRoll[1] = "Sumit"                              //adding key and value
	studentRoll[2] = "Amit"                               //adding key and value
	studentRoll[3] = "Rahul"                              //adding key and value
	fmt.Println("studentRoll map contents:", studentRoll) //displaying contents of map

	/*
			you can also initialise the map during declaration :
			studentRoll := map[int]string {
		        1: "Sumit",
		        2: "Amit",
		    }
	*/
}

/*
Output :

studentRoll map contents: map[1:Sumit 2:Amit 3:Rahul]

*/

/*
Note : Try to add values with same key and then find the output. Also try to add same values with different keys.
*/
