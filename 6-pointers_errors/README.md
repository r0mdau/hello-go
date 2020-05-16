#### pointers
Need to pass a reference when using very large data or to have only one instance of database connection pool for example.

#### nil
When a function returns a pointer to something, need to check if it's nil or raise a runtime exception, 
compiler will not help.

#### errors
Errors are the way to signify failure.

Don't just check errors, handle them gracefully :
https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

#### new types
useful for adding more domain specific meaning to values