/*
delete(map, key) is the syntax to delete key from a map.
The delete function does no return any value.
*/

package main

import (
	"fmt"
)

func main() {
	studentRoll := map[int]string{
		1: "Sumit",
		2: "Amit",
		3: "Rahul",
	}
	fmt.Println("map before deletion", studentRoll)
	delete(studentRoll, 3)
	fmt.Println("map after deletion", studentRoll)

}

/*
Output :

map before deletion map[2:Amit 3:Rahul 1:Sumit]
map after deletion map[1:Sumit 2:Amit]

*/
