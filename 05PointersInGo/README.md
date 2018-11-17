# Pointer  

In computer science, a pointer is a programming language object that stores the memory address of another value located in computer memory. A pointer references a location in memory, and obtaining the value stored at that location is known as dereferencing the pointer.  

As an analogy, a page number in a book's index could be considered a pointer to the corresponding page; dereferencing such a pointer would be done by flipping to the page with the given page number and reading the text found on that page.  

For example :  

![alt text](https://cdncontribute.geeksforgeeks.org/wp-content/uploads/pointers-in-c.png)  

In the above example variable var having value 10 is stored in the memory at a particular address and this address is stored by a pointer prt.  

[Learn more about pointers](https://en.wikipedia.org/wiki/Pointer_(computer_programming) "More on pointers")  

## Declaration of pointers in Go  

** *T ** is the type of the pointer variable which points to a value of type T.  
The ** & ** operator is used to get the address of a variable. 