# learn-golang


## go.mod file
In a React project, we download packages via _npm_ and import them in the components. Those downloaded packages are stored in `node_modules`.

In a Golang project, when we import packages contained in other modules, we need to manage them in a separate file called `go.mod`. It is a module file for our own project.

**Create a go.mod file :-**

```bash
go mod init <file_name>
```

Note: The convention is to use the _repository name_ for the `go.mod` file in case you plan to publish your module for others to use. This helps _Go tools_ to easily track and download your module.

## Run a golang code

**Step 1 :** In the terminal, navigate to the _folder_ with your golang code.
**Step 2 :** Enter below command -:

```shell
go run .
```