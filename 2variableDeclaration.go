package main

import(
	"fmt"
)

func main(){
	var x int = 5
	fmt.Println(x) //when no value is provided to x,then 0 is printed

	y := 2 //variables can also be declared in this way without the 'type' because Go can infer it
	var sum int = x + y
	fmt.Println(sum)
}
