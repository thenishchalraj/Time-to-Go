package main

import(
	"fmt"
	"errors"
	"math"
)

func main(){

	result1 := sum(3, 7)
	fmt.Println(result1)

	result2, err := sqrt(25) //you can change the argument to check for the negative number, also Go doesn't has exception
	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(result2)
	}
}

func sum(x int, y int) int {
	return x + y
}

//in Go, functions have multiple return values
func sqrt(z float64) (float64, error){
	if z < 0 {
		return 0, errors.New("undefined for negative number")
	}
	return math.Sqrt(z), nil //we don't have to return error here hence we will return nil
}
