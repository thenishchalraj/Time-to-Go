package main

import(
	"fmt"
)

func main(){
	var x int
	var y int
	fmt.Println("Enter first number")
	fmt.Scan(&x)
	fmt.Println("Enter second number")
	fmt.Scan(&y)

	fmt.Println("and the sum is", x+y)
}
