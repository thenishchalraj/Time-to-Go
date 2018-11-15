/*
Like we have anonymous structure, we can have anonymous fields in the structure i.e. fields without any name
*/
package main

import (
	"fmt"
)

type Student struct {
	string //anonymous field
	int
}

func main() {
	stud1 := Student{"Sumit", 20}
	fmt.Println(stud1)

	/*
		Even though an anonymous fields does not have a name, by default the name of a anonymous field is the name of its type.
		For example in the case of the Student struct above, although the fields are anonymous, by default they take the name of the type of the fields.
		So Student struct has 2 fields with name string and int.
	*/
	var stud2 Student
	stud2.string = "Amit" //here name of field is taken as string
	stud2.int = 20
	fmt.Println(stud2)
}

/*
Output :

{Sumit 20}
{Amit 20}
*/
/*
Note : If you are having two fields of string type or any other type but same and you are trying to assign
value to it then it will give error as :
# command-line-arguments
./4anonymousFieldsInGo.go:12:2: duplicate field string

So try to name them.(there may be some other process to do this. I will update if some other process is there )

type Student struct {
	string
	string //error
	int
}
*/
