# Dependency tracking for your code

```NOTE```
Go takes an unusual approach to dependency management, in that it is source-based instead of
artifact-based. In an artifact-based dependency management system, packages consist of artifacts
generated from source code and are stored in a separate repository system from source code.


When your code imports packages contained in other modules, you manage those dependencies  through
your code's own module. That module is defined by a go.mod file that tracks the modules that provide 
those packages. That go.mod file stays with your code, including your source code repository

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, 
giving it the name of the module your code will be in.

In actual development, the module path will typically be the repository location where your source code 
will be kept. For example, the module path might be github.com/mymodule. if you plan to publish your module 
for others to use, the module path must be a location from which Go tools can download your module.

For learning purposes, we run the command that is written below
```bash 
 $ go mod init example/hello
 go: creating new go.mod: module example/hello
```
And we created hello.go file. You can check the code out from the file

After writing our logic to our file. We must run to see the result. In order to run the code write this

```bash
$ go run .
here you can see the result! 
```

