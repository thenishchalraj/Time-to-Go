/*
Like functions, methods can also be defered.
*/
package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
}

func (s student) fullName() {
	fmt.Printf("%s %s", s.firstName, s.lastName)
}

func main() {
	s := student{
		firstName: "Sumit",
		lastName:  "Mishra",
	}
	defer s.fullName() //will be executed before return of main() function
	fmt.Printf("Welcome ")
}

/*
Output :
Welcome Sumit Mishra
*/
