# Leanred from This Section

This README will document things learned from this section.

## Iterations

Go has no concept for a while loop.  Neither `do` nor `until` are keywords in language.

The one looping mechanism that exists is the `for` loop.  The composition of the `for` loop in go is similar to other languages:

```go
for <initilized var>; <exit criteria>; <var modification> {
    <do stuff here>
} 
```

In addition to the standard go tests, go has (benchmarking)[https://pkg.go.dev/testing#hdr-Benchmarks] built in.

## Go Tips and Tricks

The `:=` operator is shorthand to both initialize a variable and set it equal to something.

So this:

```go
testvar := "this is a test string"
```

Is functionally equal to this:

```go
var testvar string
testvar = "this is a test string"
```

Go has verbose (package documentation)[https://pkg.go.dev/].