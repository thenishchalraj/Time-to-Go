package main //always the first line goes the package name, main is the package that helps this language in execution

import(
	"fmt"  //this is formatted IO and is used to provide the standard input and output in Go, just like printf and scanf in C
		//we can import as many libraries as we want to import from the package
)

func main(){ //func keyword followed by the name of the function
//main function is same as that in c++ and java. It is the function that will be executed automatically when you run the programm.
	fmt.Println("Hello World") //Simply print the traditional "Hello World" line on the console
}
