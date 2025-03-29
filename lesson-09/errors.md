# Errors ⚠️
Errors in Go are handled **explicitly** using the built-in `error` type. An Error is any type that implements the simple built-in **error interface**.

```go
type error interface {
    Error() string
}
```

- `Error()` method returns a _string_ describing what went wrong.

Go encourages error handling using **return values**. When something can go wrong in a function, that function should return an _error_ as its **last** return value.

The returned _error_ is then handled by checking whether the `error` is `nil`

```go
 i, err := strconv.Atoi("42b")
 if err != nil {
    fmt.Println("counld not convert", err)
    return
 }

 fmt.Println("converted value is", i)
```
- Since, `err` is a variable of `error` interface, the guard clause checks whether `err` is _not nil_, meaning there is an error.

## :sparkles: Creating and Returning an Error
Functions return errors as their _last_ values. 

```go
package main

import (
    "errors"
    "fmt"
)

// Function that returns an error
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Error: cannot divide by zero
    } else {
        fmt.Println("Result:", result)
    }
}
```
- Here, `divide()` returns 2 values &mdash; float64 and error
- If `b` is 0, we return an _error_ using `errors.New()`
- `main()` checks for `err` explicitly

## :sparkles: Defining Custom Error Types
Since `error` is an interface with a single method `Error()` that returns a string, we can use this to create **custom error structs** to hold fields for _detailed_ error messages.

```go
type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func test() error {
    return &MyError{404, "Not Found"}
}

func main() {
    err := test()
    fmt.Println(err) // Error 404: Not Found
}
```