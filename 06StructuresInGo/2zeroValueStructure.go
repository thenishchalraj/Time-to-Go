/*
When a struct is defined and it is not explicitly initialised with any value, the fields of the struct are assigned their zero values by default.
*/

package main

import (
	"fmt"
)

type Student struct {
	firstName, lastName string
}

func main() {
	var stud1 Student //zero valued structure
	stud2 := Student{ //assigning some value
		firstName: "Sumit",
	}
	fmt.Println("Student 1", stud1)
	fmt.Println("Student 2", stud2)
}

/*
Output :
Student 1 { }
Student 2 {Sumit }

*/
