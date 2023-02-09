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

Slices in go do NOT encode the size of the collection, they can be of any size.

## Go Tips

Use `range` to traverse arrays and slices.

Theres a go package called `reflect` which has a `DeepEqual` that is better suited to test equality amongst slices but it is not type safe.  That is to say, it will compile even if it's comparing a string and a slice but it wont run.