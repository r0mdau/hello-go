#### to know

The key in map can only be a comparable type : https://golang.org/ref/spec#Comparison_operators

But the value type can be everything.

Never initialize an empty map variable, but use make :
```go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```