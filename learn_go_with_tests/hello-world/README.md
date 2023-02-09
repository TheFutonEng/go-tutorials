# Leanred from This Section

This doc will just cover high level things that I learned.

## Modules

I still do not understand how modules work.  Modules were disabled on my system after trying to use Go workspaces.

This will happen anywhere in the directory structure when trying to create a module. 

```
[rmengert@Robs-MBP:~]
$ go mod init got-test
go: modules disabled by GO111MODULE=off; see 'go help modules'
[rmengert@Robs-MBP:~]
```

This doesn't impact being able to proceed through the material and run `go test` where needed.  

#TODO: Understand modules and fix my laptop to be able to use them properly.

## Testing

Testing in Go is built directly into the language.  Writing a test is basically writing a function with a couple of rules:

- The name of the file container the test functions must be of the format `\<someIdentifier>_test.go`
- The name of the test function must start with "Test"
- Test functions only take one argument of `t *testing.T`
- In order to use the above type, `testing` must be imported.

Subtests are used to test different test cases for a particular function.

```go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
```

Each `t.Run` is targeting a different test case.

## Flow Control

In Go, `if` statement operate much as they do in other languages.  `switch` statements act on a variable.  For instance:

```go
switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
```