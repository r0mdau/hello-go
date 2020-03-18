Remember that slices of array points to the same memory address
```go
func main() {
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	
	y := x[:] // slice "y" points to the underlying array "x"
	
	z := make([]string, len(x))
	copy(z, x[:]) // slice "z" is a copy of the slice created from array "x"

	y[1] = "Belka" // the value at index 1 is now "Belka" for both "y" and "x"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
}
//--- prints

[3]string [Лайка Belka Стрелка]
[]string [Лайка Belka Стрелка]
[]string [Лайка Белка Стрелка]
```

Remember to copy slices of array in other variable to garbage collect original
```go
func main() {
	a := make([]int, 1e6) // slice "a" with len = 1 million
	b := a[:2] // even though "b" len = 2, it points to the same the underlying array "a" points to
	
	c := make([]int, len(b)) // create a copy of the slice so "a" can be garbage collected
	copy(c, b)
	fmt.Println(c)
}
```