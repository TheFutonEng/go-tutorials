# Leanred from This Section

Need to review Structs, Methods and interfaces again.  Yes, I learned that in this section.

Pointers are used to point at a location in memory.

## Goisms

In go, when you call a function or a method, the arguments are copied.  For passing strings, ints, etc, this is no big deal.  But if we're operating against a bank account, modifying the copy doesn't modify the realy thing and that's where pointers are needed.

Struct Pointers are automatically dereferenced.

Go let's you create types (structs) from existing types:

```go
type Bitcoin int

type Wallet struct {
    balance Bitcoin
}
```

There's some type of rule around first letter capitalization when it comes to methods and functions

## Syntax Discoveries

Pointers are specified in function or methods by putting an `*` in front of the receiver type.

Regular method:

```go
func (w Wallet) Deposit(amount int) {}
```

Method that takes a pointer to a Wallet as a receiver type:
```go
func (w *Wallet) Deposit(amount int) {}
```