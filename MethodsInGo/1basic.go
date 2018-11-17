/*
This is a basic example of methods in Go
*/
package main
import (
	"fmt"
)
type Student struct {
	name   string
	branch string
}
/*
 displayStudent() method has Student as the receiver type
*/
func (s Student) displayStudent() {
	fmt.Println("Name :" + s.name)
	fmt.Println("Branch :" + s.branch)
}
func main() {
	stud1 := Student{
		name:   "Sumit Mishra",
		branch: "CSE",
	}
	stud1.displayStudent() //Calling displayStudent() method of Student type
}
/*
Output :
Name :Sumit Mishra
Branch :CSE
*/
