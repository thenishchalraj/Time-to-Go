/*
Maps are reference types i.e. if we assign a map to another then both of the maps will point to the same
location. Changes in one map will reflet the changes in the other and vice-versa.
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
	fmt.Println("Student roll salary", studentRoll)  //print original map
	newStudentRoll := studentRoll                    //assigning one map to other
	newStudentRoll[4] = "Alok"                       //changing in new map
	fmt.Println("Student roll changed", studentRoll) //printing old map

}

/*
Output :

Student roll salary map[1:Sumit 2:Amit 3:Rahul]
Student roll changed map[1:Sumit 2:Amit 3:Rahul 4:Alok]

*/
