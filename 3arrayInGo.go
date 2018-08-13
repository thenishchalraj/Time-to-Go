package main

import(
	"fmt"
)

func main(){
	var array [5]int //this array holds 5 integers
	fmt.Println(array) //to print the whole array, when nothing is stored in the array then it will output only 0 elements

	//lets input an integer at place 2nd
	array[2] = 7
	fmt.Println(array)

	//declare the elements at a time
	arr := [5]int{5,6,4,3,2} //don't use the same variable name here for the array as it has once been initialised and now will give error
	fmt.Println(arr)

	//creating a sliced array i.e. this will remove the commas from in between the elements
	brr := []int{5,4,3,2,1} //just remove the number of elements to be there and then you can append as many as you want
	brr = append(brr,0)
	fmt.Println(brr)
}
