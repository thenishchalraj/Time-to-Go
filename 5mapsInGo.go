package main

import(
	"fmt"
)

func main(){
	//maps are like dictionary in python and associates in php i.e. with key & value pair
	vertices := make(map[string]int)

	vertices["number1 "] = 1
	vertices["number2 "] = 2
	vertices["number3 "] = 3

	fmt.Println(vertices) //will output the whole map

	fmt.Println(vertices["number 2"]) //will output 0 as it can't find the value for the key given
	fmt.Println(vertices["number2 "]) //will output the particular value of the key given

	delete(vertices,"number1 ") //use the delete function to delete something from the map
	fmt.Println(vertices)
}
