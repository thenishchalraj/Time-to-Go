/*
We have learned about anonymous fields. But what if that field is a structure itself. For example :

type Name struct {
	firstname, lastname string
}
type Student struct {
	Name //anonymous field of type struct
	age    int
	branch string
}

In the above code snippet, the Student struct has an anonymous field Name which is a struct.
Now the fields of the Name struct namely firstname and lastname are called promoted fields
since they can be accessed as if they are directly declared in the Student struct itself.
*/
package main

import (
	"fmt"
)

type Name struct {
	firstname, lastname string
}
type Student struct {
	Name   //anonymous Field
	age    int
	branch string
}

func main() {
	var stud1 Student
	stud1.Name = Name{
		firstname: "Sumit",
		lastname:  "Mishra",
	}
	stud1.age = 20
	stud1.branch = "CSE"
	fmt.Println("Name:", stud1.firstname+" "+stud1.lastname) //promoted fields
	fmt.Println("Age:", stud1.age)
	fmt.Println("Branch:", stud1.branch)
}

/*
Output :

Name: Sumit Mishra
Age: 20
Branch: CSE
*/
