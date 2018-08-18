package main

import(
	"fmt"
)

func main(){
	var x int //first number
	var y int //second number
	var z int //option can be 1,2,3,4 for D,M,A,S

	//let's get the numbers stored in the variables
	fmt.Println("Enter first number")
	fmt.Scan(&x) //taking from input
	fmt.Println("Enter second number")
	fmt.Scan(&y)

	//ask for what to do with the numbers
	fmt.Println("Select from the options available below: \n 1. Division \n 2. Multiplication \n 3. Addition \n 4. Substraction")
	fmt.Scan(&z)
	switch z {
		case 1:
			fmt.Println("Result of division: ", d(x,y)) //calling functions for each operation
		case 2:
			fmt.Println("Result of multiplication: ", m(x,y))
		case 3:
			fmt.Println("Result of division: ", a(x,y))
		case 4:
			fmt.Println("Result of division: ", s(x,y))	
	}
}
 
//functions for the operations
func d(a int, b int) int {
	return a/b
}

func m(a int, b int) int {
	return a*b
}

func a(a int, b int) int {
	return a+b
}

func s(a int, b int) int {
	return a-b
}
