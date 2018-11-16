package main //always the first line goes the package name, main is the package that helps this language in execution

import(
	"fmt"  //this is formatted IO and is used to provide the standard input and output in Go, just like printf and scanf in C
//we can import as many libraries as we want to import from the package
)

func twoNumberSum(a int, b int) int{ //func keyword followed by the function name, here parameters are given same as that in c++ and java
//only difference is the return type. In go, you write the return type after the bracket i.e. in the example 'int' is the return type
	return a + b
}

func threeNumberSum(a, b, c int) int { //yeah, something cool is here. If you have same type of variable then you can 'go' with writing the 
//variable type just after the names of all variable
    return a + b + c
}

func main(){ //func keyword followed by the name of the function
//main function is same as that in c++ and java. It is the function that will be executed automatically when you run the programm.

	fmt.Println("Hello World") //Simply print the traditional "Hello World" line on the console

	res := twoNumberSum(1, 2) //declare a variable res (don't worry you will learn more about variables later) and store the result obtained from
//calling twoNumberSum() function with (1,2) as parameter	 
    fmt.Println("1 + 2 = ", res) // printing the result

    res = threeNumberSum(1, 2, 3) // yeah you can use the same variable to store results of two different function but the previous value will be lost.
    fmt.Println("1 + 2 + 3 = ", res)
}

/*
compile the go file by going into the directory of file and type 'go build 1basicFunction.go'
run the go file by 'go run 1basicFunction.go'
other ways of compiling and running is simply go the file's location(/go/src/yourProject) and use 'go build' to compile your program and 'go run' to run your program.
*/


/*
output

Hello World
1 + 2 = 3
1 + 2 + 3 = 6

*/