## What is Defer ?  

Defer statement is used to execute a function call just before the function where the defer statement is present returns i.e. if a defer statement is present in a function then that differ statement will be executed just before the return statement of the function irrespective of the place of it's declaration.  
This is useful for many reasons, the most common of which is to close an open connection immediately before the function ends.  

Another reason may be, it's not always guaranteed that your code may reach the end of the function (e.g. an error or some other condition may force you to return well ahead of the end of a function). The defer statement makes sure that whatever function is assigned to it gets executed for sure even if the function panics or the code returns well before the end of the function.  

The defer statement also helps keep the code clean esp. in cases when there are multiple return statements in a function esp. when one needs to free resources before return (e.g. imagine you have an open call for accessing a resource at the beginning of the function - for which a corresponding close must be called before the function returns for avoiding a resource leak. And say your function has multiple return statements, maybe for different conditions including error checking. In such a case, without defer, you normally would call close for that resource before each return statement). The defer statement makes sure the function you pass to it is always called irrespective of where the function returns, and thus saves you from extraenous housekeeping work.  

**NOTE :** When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.