package main

import "fmt"

/*
We use equality operators (" == ") to check if two items are equal or not.
But in maps equality operators can't check the equality of 2 maps.package MapsInGo
"==" can only be used to check if both the maps are nil
*/

func main() {
	studentRoll := map[int]string{
		1: "Sumit",
		2: "Amit",
	}

	newStudentRoll := studentRoll

	if studentRoll == newStudentRoll {
		fmt.Println("Both maps are equal")
	}
}

/*
Output :

# command-line-arguments
./5mapsEquality.go:19:17: invalid operation: studentRoll == newStudentRoll (map can only be compared to nil)
*/
