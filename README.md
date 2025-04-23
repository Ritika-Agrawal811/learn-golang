# learn-golang

## :sparkles: Why learn Go?
- Go is a _compiled_ language. It has _greater_ **execution speed** as compared to interpreted languages like JavaScript, PHP, Python and Ruby.

- In comparison to other compiled languages like C/C++, Java, Rust and C#, Go has _greater_ **compilation  speed**.

- Go is _lightweight_, it has a small **runtime** compared to Java and does not require a Virtual Machine to run.

- It includes **garbage collection** that is designed to be efficient with low pause times.

## :sparkles: go.mod file
In a React project, we download packages via _npm_ and import them in the components. Those downloaded packages are stored in `node_modules` and maintained in `package.json`.

Similarly, in a Golang project, when we _import_ packages contained in other modules, we need to manage them in a separate file called `go.mod`. It is a **module file** for our own project.

#### Create a go.mod file 

```bash
go mod init <file_name>
```

Note: The convention is to use the _repository name_ for the `go.mod` file in case you plan to publish your module for others to use. This helps _Go tools_ to easily track and download your module.

## :sparkles: Run a golang code

#### Step 1 : Navigate to Your Go Project Folder  
Open a terminal and move to the directory where your Go project is located:

```shell
cd /path/to/your/project
```

#### Step 2 : Run the Go Code

Use the following command to _execute_ your Go program:
```shell
go run <path to main package>
```
or, if you're already inside the main package directory:

```shell
go run .
```

This command _compiles_ and runs the Go files without generating a **binary**.