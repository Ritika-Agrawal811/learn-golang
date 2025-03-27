# Functions 🛠️
Functions are _resuable_ blocks of code that perform a certain **task**. They help in code organization and readability.

## :sparkles: Defining a function
A function in Go is defined using the `func` keyword, followed by:

- The function name
- A parameter list (optional)
- A return type (optional)
- A function body (code to execute)


```go
// function without parameters and return value
func greet() {
    fmt.Println("Hello, Go!")
}

// function with parameters and return value
func getFullName(firstname string, lastname string) string {
    return firstname + " " + lastname
}
```

:warning: **Note:** In Go, the **parameter type** is defined _after_ the variable name.

If multiple parameters have a same _type_, we can also group them like this :point_down:

```go
func getFullName(firstname, lastname string) string {
    return firstname + " " + lastname
}
```

when a function has multiple return values, we need to surround them in `parenthesis ()` :point_down:

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}
```

## :sparkles: Named Return Values
Go allows named return values, which act as **variables** inside the function.

```go
func rectangleArea(length, width int) (area int) {
    area = length * width // `area` is already declared
    return // No need to specify `area` explicitly
}
```

- **return variable** is declared in the _function signature_.
- No need to write `return area`, just `return` works.

:warning: **Note:** It is always better to mention the returned variables explicitly for verbose. This helps to avoid unnecessary confusion.

## :sparkles: Guard Clause
An early return from a function using conditionals is called a **guard clause**.

This code :point_down:

```go
func checkAccess(user string, hasPermission bool) {
    if user != "" {
        if hasPermission {
            fmt.Println("Access granted")
        } else {
            fmt.Println("Access denied")
        }
    } else {
        fmt.Println("Invalid user")
    }
}
```

can be simplified as :point_down:

```go
func checkAccess(user string, hasPermission bool) {
    if user == "" {
        fmt.Println("Invalid user")
        return
    }

    if !hasPermission {
        fmt.Println("Access denied")
        return
    }

    fmt.Println("Access granted")
}
```

<p align="center">· · · · · · · · · · · · ·</p>

# Blank Identifier 🚫
Go doesn't allow us to have **unused variables**, which is a good feature because it avoids any unecssary memory consumption. 

:thought_balloon: So how can we ignore variables returned from a function that we don't want to use anywhere?

Go provides a **blank identifier ( _ )**. It is a special identifier used to ignore a value.  It acts as a **placeholder** for values we _don’t want_ to use in a program.

#### a) Ignoring Returned Values

A function may return more than one values. We can use `blank identifier (_)` to ignore the ones we don't need.

```go
 q, _ := divide(10, 3) // Ignoring remainder
```
Here we are _ignoring_ the second returned value from `divide` function.

#### b) Ignoring Unused Imports

Go does not allow unused imports. If we need to import a package only for its _side effects_, we can use `_` to suppress the error.

```go
import _ "database/sql"
```


