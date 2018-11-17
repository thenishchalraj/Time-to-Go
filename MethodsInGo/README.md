# Methods In Go  
A method is just a function with a special receiver type that is written between the func keyword and the method name. The receiver can be either struct type or non struct type. The receiver is available for access inside the method.  
Syntax :  
```go
func (t Type) methodName(parameter list) {  
}
```  
### Why methods?  
We know that we can solve the same problem with the help of functions as well as methods. So, why do we need methods while functions are already there.  
The reason behind the this may be :  
  * [Go is not a pure object oriented programming language](https://golang.org/doc/faq#Is_Go_an_object-oriented_language) and it does not support classes. Hence methods on types is a way to achieve behaviour similar to classes.
  
  * Methods with same name can be defined on different types whereas functions with the same names are not allowed. Lets assume that we have a Square and Circle structure. It's possible to define a method named Area on both Square and Circle. (see example **2squareAndCircle.go**)
