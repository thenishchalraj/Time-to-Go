/*
If you want to change the contents of an array through a function then don't pass pointer of array as an argument to the function.
No error will come if you do so. But the desired thing can be done even without using pointer i.e. by using slice of array.
*/

package main

import (
	"fmt"
)

func modifyArray(slice []int) {
	slice[0] = 2
}

func main() {
	a := [3]int{1, 2, 3}
	modifyArray(a[:]) //passing slice of array a
	fmt.Println(a)
}

/*
Pointer way of doing the same :
package main

import (
    "fmt"
)

func modifyArray(arr *[3]int) {
    arr[0] = 2
}

func main() {
    a := [3]int{1, 2, 3}
    modifyArray(&a)
    fmt.Println(a)
}
*/

/*
Output :

[2 2 3]
*/
