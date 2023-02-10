# Leanred from This Section

Arrays in Go have a fixed capacity which seems pretty limiting until you learn that slices are a thing.  This is to say, if a function is expecting an array of size five and gets passed an array of size 6, the program will not compile

The [`range`](https://gobyexample.com/range) command assists in traversing arrays and slices in a more organized way.  For each iteration, range returns two values: the index and the actual value in the array.

```go
numbers := []int{1,2}
for idx, number := range numbers{
    <do some stuff>
}
```

The `range` command also supports the ability to basically ignore either the index or the value through the blank identifier `_`.  That same loop with the index ignored:

```go
numbers := [2]int{1,2}
for _, number := range numbers{
    <do some stuff>
}
```

## Slices

Slices in go do NOT encode the size of the collection, they can be of any size.  However slices do have a capacity.  A slice of size two can't have an element inserted into the 10th position.  However, the append function can take a slice and a new value and combine them into a new slice.

## Common Practices

If a slice is create from an array or another slice, the slice references the original object in memory.  This means that is a function returns a slice of an array, it's better off copying the return value to a new slice so the reference to the original array is released and can be garbage collected.  Example [here](https://go.dev/play/p/Poth8JS28sc).

## Go Tips

Use `range` to traverse arrays and slices.

Theres a go package called `reflect` which has a `DeepEqual` that is better suited to test equality amongst slices but it is not type safe.  That is to say, it will compile even if it's comparing a string and a slice but it wont run.

A [variadic function](https://gobyexample.com/variadic-functions) is one which will take a variable number of arguments.  The `...` in the function defintion is what allows for this.

```go
func SumAll(numbersToSum ...[]int) []int {
	return nil
}
```