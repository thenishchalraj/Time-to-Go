# Structures in Go  

A structure is a user defined type which represents a collection of fields. It can be used in places where it makes sense to group the data into a single unit rather than maintaining each of them as separate types.  

For instance a student has a firstName, lastName, age, enrollment_number, branch, year. It makes sense to group these six properties into a single structure student.  

### Declaring a Structure  

```go
type Student struct {  
    firstName string
    lastName  string
    age       int
    enroll    string
    branch    string
    year      int
}
```  

The above snippet declares a structure type Student which has fields firstName, lastName, age, enroll, branch and year. This structure can also be made more compact by declaring fields that belong to a same type in a single line followed by the type name. For examplee :  

```go
type Student struct {  
    firstName, lastName, enroll, branch string
    age, year                           int
}
```  

The above Student struct is called a *named structure* because it creates a new type named Student which can be used to create structures of type Student.  

### Anonymous Structures  

It is possible to declare structures without declaring a new type and these type of structures are called *anonymous structures*.  

```go
var student struct {  
        firstName, lastName, enroll, branch string
        age, year int
}
```  

The above snippet creates a anonymous structure student.