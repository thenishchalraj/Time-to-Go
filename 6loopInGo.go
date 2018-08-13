package main

import(
	"fmt"
)

func main(){
	//the only type of loop in Go is for loop
	for i := 0; i < 6; i++ { //everything is same, just remove the brackets
		fmt.Println(i)
	}

	//for loop also works as the while loop
	j := 0
	for j < 6{
		fmt.Println(j) //so this will return the same numbers as the above loop
		j++
	}
}
