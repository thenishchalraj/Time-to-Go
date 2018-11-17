/*
Anonymous functions are those functions having no name but arguments and return types are there.
Anonymous functions are useful when you want to define a function inline without having to name it.
*/

package main

import "fmt"

func funcCount() func() { //this function will return the number of times func() will be called
	n := 0
	return func() { //anonymous Function
		n++
		fmt.Println(n)
	}
}

func main() {
	f1, f2 := funcCount(), funcCount() //calling funcCount() two times
	f1()                               // 1<---
	f2()                               // 1   | <---
	f1()                               // 2<---    |
	f2()                               // 2<-------
}

/*
Output :
1
1
2
2
*/
