package main

import(
	"fmt"
)

func main(){
	x := 5

	if x < 6 { //we don't use the brackets here in Go
		fmt.Println("x is smaller than 6")
	}else if x > 2 { //this is how else if is used
		fmt.Println("x is greater than 2")
	}else { //it is necessary to use blocks for single else
		fmt.Println("x is :)")
	}
}
