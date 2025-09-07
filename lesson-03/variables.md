# Variables 🔠 
Variables in Go store values of _different_ data types. Since Go is **statically typed**, each variable has a _fixed type_ that is determined at compile time.

## :sparkles: Declaring variables
Go provides multiple ways to declare a variable

### a) Using `var` keyword with explicit type

```go
var name string = "GoLang"
var age int = 25
var pi float64 = 3.1415
var isReady bool = true
```

- variable type is _explicity_ mentioned. Helpful to use for complex types.

### b) Using `var` keyword without type

```go
var city = "New York" // Go infers type as string
var count = 100       // Go infers type as int
```

- Go automatically infers the type based on the assigned value.

**Type inference** means that the compiler automatically determines the type of a variable based on its assigned value.

### c) Short Declaration: Using `walrus` operator

```go
age := 23
```
- `walrus` operator (:=) declares and assigns a variable in the same line.
- Go _infers_ the type from the assigned value
- `walrus`operator can be used only inside _functions_.

### d)  Multiple Variable Declaration 
Go allows us to declare multiple variables in 2 ways →

#### i) **In single line :**

```go
var a, b, c int = 1, 2, 3
x, y, z := "Go", 3.14, true
```

#### ii) **group them using `var (...)` :**

```go
var (
    username string = "admin"
    password string = "1234"
    port     int    = 8080
)
```

<p align="center">· · · · · · · · · · · · ·</p>

## :sparkles: Constants

Constants are declared with the `const` keyword. They can not be declared via short declaration syntax &mdash; `walrus` operator.

As their name says, they are **fixed values**. We _can't reassign_ these variables with new values.

```go
const Pi = 3.1415
const AppName = "GoApp"
```

### Multiple Constants Declaration 📜

#### i) **In single line :**

```go
const Pi, Language, Version = 3.1415, "Go", 1.21
```

#### ii) **group them using `const (...)` :**

```go
const (
    Pi       = 3.1415
    Language = "Go"
    Version  = 1.21
)
```

- Constants can only be simple _primitve types_ &mdash; integers, boolean, strings and floats. They **can not** be slices, maps and structs.

- Constants must be known at compile time. So they usually have static values. Go allows **computed constants** as long as it can be computed at _compile time_.

```go
const firstname = "Ritika"
const lastname = "Agrawal"
const fullname = firstname + " " + lastname
```

- Computed constants can **only** be computed using other `const` variables or constants.

