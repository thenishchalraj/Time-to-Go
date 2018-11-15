/*
Structures are compareable if each of their fields are comparable.
Two struct variables are considered equal if their corresponding fields are equal.

Struct variables are not comparable if they contain fields which are not comparable.

For example : strings and int are compareable whereas maps are not compareable. So a structure containing
map is not not copareable while a structure containing string is compareable
*/

package main

import (
	"fmt"
)

type Student struct {
	firstName string
	lastName  string
}

type StudentData struct {
	data map[int]string
}

func main() {
	stud1 := Student{"Sumit", "Mishra"}
	stud2 := Student{"Sumit", "Mishra"}
	if stud1 == stud2 {
		fmt.Println("Student 1 and Student 2 are equal")
	} else {
		fmt.Println("Student 1 and Student 2 are not equal")
	}

	stud3 := Student{"Amit", "Mishra"}
	stud4 := Student{}
	stud4.firstName = "Rahul"
	if stud3 == stud4 {
		fmt.Println("Student 3 and Student 4 are equal")
	} else {
		fmt.Println("Student 3 and Student 4 are not equal")
	}

	/*
		stud1 := StudentData{data: map[int]string{
			1: "Sumit",
		}}
		stud2 := StudentData{data: map[int]string{
			1: "Sumit",
		}}
		if stud1 == stud2 {
			fmt.Println("Student 1 and Student 2 are equal")
		}
	*/
}

/*
Output :

Student 1 and Student 2 are equal
Student 3 and Student 4 are not equal

The commented part will give error as comparision is not possible between maps. The error is :

# command-line-arguments
./7structureEquality.go:50:12: invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)

*/
