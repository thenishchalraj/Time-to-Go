/*
Variadic Functions are functions which can be called with any number of trailing arguments. For example, fmt.Println() is a common variadic function.
Learn more about Variadic Function from : https://en.wikipedia.org/wiki/Variadic_function
Learn about Variadic Function in c from : https://www.geeksforgeeks.org/variadic-function-templates-c/
*/

package main

import "fmt"

func multiplyNumbers(nos ...int) { //Function that takes a number of int inputs (here number of arguments are not defined)
	fmt.Print(nos, " ") //printing all the arguments
	mult := 1
	for _, num := range nos {
		mult *= num
	}
	fmt.Println(mult) //printing the result of multiplyNumbers() function
}

func main() {
	multiplyNumbers(11, 12)     //passing 2 arguments
	multiplyNumbers(11, 12, 13) //passing 3 arguments

	nums := []int{11, 12, 13, 14} //declaring array of type 'int'
	multiplyNumbers(nums...)      //passing 4 argumnets as array
}

/*
Output :
[11 12] 132
[11 12 13] 1716
[11 12 13 14] 24024
*/
