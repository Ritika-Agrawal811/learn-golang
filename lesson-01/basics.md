# Package Declaration

A package is a way to _group together_ related Go source files to organize and structure code. It is made up of _all_ the files in the _same_ directory.

Every `.go` file must start with a package declaration. 

```go
package main
```

## Types of Packages

Go has **two** main types of packages:

1. `main` Package
    - The main package is _special_ in Go.
    - It is the _entry point_ of a Go program.

Programs **must** have a `main` package with a `main()` function. This `main` function is executed first when a go program is run.

Example :
```go
    package main

    import "fmt"

    func main() {
        fmt.Println("Hello, World!")
    }
```

2. Library Packages
    - These are reusable and do not need a `main()` function.
    - They are used to share functions, logic etc.

Example :
```go
    package utils

    func Add(a int, b int) int {
        return a + b
    }
```

# `fmt` package

`fmt` is a _standard_ package included in golang installation. It contains functions for formatting texts, including concatenation and printing to console.

# Use external packages

We can use externals packages in our go code by importing them and calling their functions.

**Step 1 :**  Visit [pkg.go.dev](https://pkg.go.dev/)
**Step 2 :**  Locate any package you wish to use. In the _Documentation_ section, you can view a package's functions
**Step 3 :** From the top of the page, copy the package name and import it in your code. 

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

**Step 4 :** Add new module requirements and sums

When we include a third-party package, we need to run `go mod tidy` to -:

- Ensure that the `go.mod` file reflects the correct dependencies used in our project.

- Add any missing modules required to build the project.

- Remove unused dependencies.

- Ensure that `go.sum` is updated with cryptographic hashes of the downloaded modules for verification

This command acts as a _cleanup_ tool for Go modules.

```shell
$ go mod tidy
go: finding module for package rsc.io/quote/v4
go: downloading rsc.io/quote v1.5.2
go: downloading rsc.io/quote/v4 v4.0.1
go: found rsc.io/quote/v4 in rsc.io/quote/v4 v4.0.1
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832
```
