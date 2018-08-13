package main

import(
	"fmt"
)

func main(){

	a := []string{"a","b","c","d"}
	//index and value pair when an array
	for index, value := range a {
		fmt.Println("index", index,"value", value) //println has a feature that we can pass multiple arguments in it
	}

	//and we can do the same thing with map
	m := make(map[string]string)
	m["a"]="alpha"
	m["b"]="beta"
	//key and value pair when a  map
	for key, value := range m{
		fmt.Println("key", key, "value", value)
	}
}
