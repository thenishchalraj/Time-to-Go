/*
Here is a basic example of structure in Go
*/
package main

import (
	"fmt"
)

type Student struct {
	firstName, lastName, enroll, branch string
	age, year                           int
}

func main() {

	//creating structure using field names
	student1 := Student{
		firstName: "Sumit",
		age:       20,
		enroll:    "0103CS16",
		lastName:  "Mishra", //order can be changed
		branch:    "CSE",
		year:      3,
	}

	//creating structure without using field names
	student2 := Student{"Amit", "Mishra", "0103CS16", "CSE", 20, 3} //same order

	fmt.Println("Student 1", student1)
	fmt.Println("Student 2", student2)
}

/*

Output :

Student 1 {Sumit Mishra 0103CS16 CSE 20 3}
Student 2 {Amit Mishra 0103CS16 CSE 20 3}
*/
