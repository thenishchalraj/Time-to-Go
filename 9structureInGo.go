package main

import (
	"fmt"
)

//structure is collection of things so you can group them together into more logical types
type person struct{
	//fields to be in the structure, each with a name and a type, it could be anything, it could also be another struct
	name string
	age int
}
 
func main(){
	p := person{name:"Nishchal Raj",age:21} //curly braces for the fields in the structure
	fmt.Println(p)
	//if you want to get the specific field then you can use the '.' (dot) notation
	fmt.Println(p.name)
}
