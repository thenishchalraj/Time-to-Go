/*
Can we do arithmetic operations on pointers in go?
Answer is NO. Go does not support pointer arithmetic which is present in other languages like C.
See how ponter arithmetic works in c at https://www.tutorialspoint.com/cprogramming/c_pointer_arithmetic.htm
*/

package main

func main() {
	a := [...]int{1, 2, 3, 4}
	ptr := &a
	ptr++
}

/*
Output :

# command-line-arguments
./4arithmeticOnPointers.go:11:5: invalid operation: ptr++ (non-numeric type *[4]int)
*/
