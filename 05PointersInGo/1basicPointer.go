package main

import (
	"fmt"
)

func main() {
	a := 10                                //declaring and initialising an int variable
	var ptr *int = &a                      //pointer pointing to a of type int
	fmt.Printf("Type of ptr is %T\n", ptr) //type of pointer
	fmt.Println("address of a is", ptr)    //address of variable
	fmt.Println("value of a is", *ptr)     //accessing value of variable
	*ptr++                                 //changing value of a
	fmt.Println("new value of a is", a)    //new value of a
}

//if you didn,t initialise the pointer then it will have a <nill> value

/*
Output :

Type of ptr is *int
address of a is 0xc0000180e0
value of a is 10
new value of a is 11

NOTE : You might get a different address for a since the location of a can be anywhere in memory.
*/
