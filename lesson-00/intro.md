# What is Go? :sparkles:

Go (also called Golang) is a modern, fast, and easy-to-use language created by Google for building scalable servers, APIs, and cloud applications.

# Why learn Go? 🤔
- Go is a _compiled_ language. It has _greater_ **execution speed** as compared to interpreted languages like JavaScript, PHP, Python and Ruby.

- In comparison to other compiled languages like C/C++, Java, Rust and C#, Go has _greater_ **compilation  speed**.

- Go is _lightweight_, it has a small **runtime** compared to Java and does not require a Virtual Machine to run.

- It includes **garbage collection** that is designed to be efficient with low pause times.

# Getting started with Go 🌱

## 📦 Creating a `go.mod` file
In a React project, we download packages via _npm_ and import them in the components. Those downloaded packages are stored in `node_modules` and maintained in `package.json`.

Similarly, in a Golang project, when we _import_ packages contained in other modules, we need to manage them in a separate file called `go.mod`. It is a **module file** for our own project.

➡️ **To create a go.mod file, use command :**

```bash
go mod init <file_name>
```

Note: The convention is to use the _repository name_ for the `go.mod` file in case you plan to publish your module for others to use. This helps _Go tools_ to easily track and download your module.

## 📝 Writing your first `main.go`

Every Go program starts from a `main` package with a `main()` function.

Create a file named _main.go_ in your project root and add this code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## :running_man: Running Go code

### Step 1 : Navigate to your Go project folder  
Open a terminal and move to the directory where your Go project is located:

```shell
cd /path/to/your/project
```

### Step 2 : Run your Go program

Use the following command to _execute_ your Go program:
```shell
go run <path to main package>
```
or, if you're already inside the main package directory:

```shell
go run .
```

This command _compiles_ and runs the Go files without generating a **binary**.