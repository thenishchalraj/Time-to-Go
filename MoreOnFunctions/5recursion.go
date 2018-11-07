/*
The process in which a function calls itself directly or indirectly is called recursion and the corresponding function is called as recursive function.
Learn more about recursion on : https://www.geeksforgeeks.org/recursion/
*/
package main

import "fmt"

func fact(n int) int { //function to print factorial (the traditional recursive function <3)
	if n == 0 {
		return 1
	}
	return n * fact(n-1) //calling function in function
}

func main() {
	fmt.Println("Factorial : ", fact(5))
}

/*
Output :
Factorial :  120
*/
