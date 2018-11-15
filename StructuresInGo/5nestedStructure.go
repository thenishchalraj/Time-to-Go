/*
It is possible that a struct contains a field which in turn is a struct.
These kind of structs are called as nested structs.
*/

package main

import (
	"fmt"
)

type Name struct {
	firstname, lastname string
}
type Student struct {
	name   Name //name of Name struct type
	age    int
	branch string
}

func main() {
	var stud1 Student
	stud1.name = Name{
		firstname: "Sumit",
		lastname:  "Mishra",
	}
	stud1.age = 20
	stud1.branch = "CSE"
	fmt.Println("Name:", stud1.name.firstname+" "+stud1.name.lastname)
	fmt.Println("Age:", stud1.age)
	fmt.Println("Branch:", stud1.branch)
}

/*
Output :

Name: Sumit Mishra
Age: 20
Branch: CSE
*/
