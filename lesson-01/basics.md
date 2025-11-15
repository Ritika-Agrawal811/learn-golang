# Package Declaration 📦

A package is a way to _group together_ related Go source files to organize and structure code. It is made up of _all_ the files in the _same_ directory.

:warning: **Note:** Every `.go` file must start with a _package declaration_.

```go
package main
```

## :sparkles: Types of Packages

Go has **two** types of packages ➜

### a) `main` package

- The main package is _special_ in Go.
- It is the _entry point_ of a Go program.

Programs **must** have a `main` package with a `main()` function. This `main` function is executed _first_ when a go program is run.

##### Example :
```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
```
### b) Library Packages (Importable package)

- These are _reusable_ and do not need a `main()` function.
- They are used to share functions, logic etc.

##### Example :
```go
    package utils

    func Add(a int, b int) int {
        return a + b
    }
```

### c) CLI Tool (Executable Package)
- These are also reusable packages but contain a main() function and use package main.
- They are designed to be run from the terminal, not imported into the Go code.
- Used for tasks like code generation, formatting, migrations, hot-reloading, etc. 

##### Examples :

- `github.com/cosmtrek/air` → hot reloader
- `golang.org/x/tools/cmd/goimports` → formatter
- `github.com/swaggo/swag/cmd/swag` → swagger tool
- `github.com/pressly/goose/cmd/goose` → migration tool


<p align="center">· · · · · · · · · · · · ·</p>

# `fmt` Package 📦

`fmt` is a _standard_ package included in golang installation. It contains functions for formatting texts, including concatenation and printing to console.

## :sparkles: Printing to Console

#### a) `Print()` function
prints strings passed as arguments _without a newline_

```go
fmt.Print("Hello")
fmt.Print("World")

/* 
  Hello World 
 */
```

#### b) `Println()` function
prints strings passed as arguments _with a newline_

```go
fmt.Println("Hello")
fmt.Println("World")

/* 
  Hello
  World 
*/
```

#### c) `Printf()` function
prints _formatted_ strings according to a **format specifier**, _without a newline_

```go
fmt.Printf("Hello %s, you are %d years old.", "John", 25)

/* 
  Hello John, you are 25 years old.
*/
```

#### Common format specifiers are :point_down:
- `%v` : Default format (works with any type)
- `%T` : Type of the value
- `%d` : Integer
- `%s` : String
- `%f` : Float
- `%.2f` : Float with 2 decimal places. Any number can be specified.
- `%t` : Boolean
- `%p` : Pointer address
- `%c` : Character
- `%q` : Quoted string/character

## :sparkles: Concatenating Strings

### a) `Sprintf()` function
returns a single string formed by concatenating the string representations of the arguments, _without adding spaces_ between arguments or a _newline_.

```go
firstName := "John"
lastName := "Doe"

info := fmt.Sprintf("Name: %s %s, Age: %d", firstName, lastName, 30)
fmt.Println(info)

/* 
  Name: John Doe, Age: 30
 */
```


<p align="center">· · · · · · · · · · · · ·</p>

# Use external packages 📦

We can use externals packages in our go code by importing them and calling their functions.

**Step 1 :**  Visit [pkg.go.dev](https://pkg.go.dev/).
**Step 2 :**  Locate any package you wish to use. In the _Documentation_ section, you can view a package's functions.
**Step 3 :** Identify the type of package and follow accordingly.

## Importing Library Packages

**Step 4 :** From the top of the page, copy the package name and import it in your code. 

Example : 

We have imported `rsc.io/quote/v4` package. It provides simple functions that return famous quotes or phrases.

```go
package main

import (
	"fmt"

	quotes "rsc.io/quote/v4"
)

func main() {
	/* 
	  Go is a function declared in "rsc.io/quote/v4" package
	*/
	fmt.Println(quotes.Go())
}
```
Here, `quotes` is an alias for the imported package. We can call it anything, but it _defaults_ to the base package name (quote) if not explicitly defined.

**Step 5 :** Add new module requirements and sums

When we include a third-party package, we need to run `go mod tidy` to :-

- Ensure that the `go.mod` file reflects the correct dependencies used in our project.

- Add any missing modules required to build the project.

- Remove unused dependencies.

- Ensure that `go.sum` is updated with cryptographic hashes of the downloaded modules for verification

This command acts as a _cleanup_ tool for Go modules.

```bash
$ go mod tidy
go: finding module for package rsc.io/quote/v4
go: downloading rsc.io/quote v1.5.2
go: downloading rsc.io/quote/v4 v4.0.1
go: found rsc.io/quote/v4 in rsc.io/quote/v4 v4.0.1
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832
```

## Installing Executable Packages

**Step 4 :** Install the package

Executable packages have a main() function and builds into an executable library, hence these must be installed using :

```bash
go install <module>@latest
```
