/*
The dot . operator is used to access the individual fields of a structure.
*/
package main

import (
	"fmt"
)

type Student struct {
	firstName, lastName string
	age                 int
}

func main() {
	stud1 := Student{"Sumit", "Mishra", 20}
	fmt.Println("First Name:", stud1.firstName) //dot(.) operator used for accessing firstname
	fmt.Println("Last Name:", stud1.lastName)
	fmt.Println("Age:", stud1.age)

	//By using dot operator you can assign the value to a zero structure
	var stud2 Student
	stud2.firstName = "Amit"
	stud2.lastName = "Mishra"
	stud2.age = 20
	//it is not compulsory to assign values to all fields. If you are not assigning then zero value will be assigned
	fmt.Println("Student 2:", stud2)
}

/*
Output :

First Name: Sumit
Last Name: Mishra
Age: 20
Student 2: {Amit Mishra 20}
*/
