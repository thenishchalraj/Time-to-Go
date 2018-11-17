/*
Here we will learn the use of methods over functions
*/
package main
import (
	"fmt"
	"math"
)
type Rectangle struct {
	length int
	width  int
}
type Circle struct {
	radius float64
}
func (r Rectangle) Area() int {
	return r.length * r.width
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Println("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Println("Area of circle %f", c.Area())
}
/*
Output :
Area of rectangle 50
Area of circle 452.389342
*/
