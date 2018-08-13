package main

import(
	"fmt"
)

func main(){
	i := 7
	fmt.Println(i) //this will output the variable's value
	fmt.Println(&i) //this will output the memory address of the variable

	inc(&i,i) //sending then address of the variable
}

func inc(x *int,y int){ //accepting the pointer by prefixing the type of variable
	*x++ //now this will de-reference the pointer
	y++ //this is incrementing the memory address
}
