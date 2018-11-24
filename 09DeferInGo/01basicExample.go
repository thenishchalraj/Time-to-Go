/*
This is a basic example of implementation of defer in go.
*/
package main

import (
	"fmt"
)

func finished() {
	fmt.Println("Finished finding largest")
}

func largest(nums []int) { //function to find the largest number of the given array
	defer finished() //will be called just before the return statement
	fmt.Println("Started finding largest")
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println("Largest number in", nums, "is", max)
}

func main() {
	nums := []int{78, 109, 2, 563, 300}
	largest(nums) //calling largest() function
}

/*
Output :

Started finding largest
Largest number in [78 109 2 563 300] is 563
Finished finding largest
*/
