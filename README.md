# Brief info about go module

```NOTE``` In a module, you collect one or more related packages for a discrete and useful set of functions.

Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies
needed to run your code, including the Go version and the set of other modules it requires.

As you add or improve functionality in your, you publish new versions of the module. Developers writing code 
that calls functions in your module can import the module's updated packages and test with the new version 
before putting it into production use.

Module is started by using the go mod init command.

Run the go mod init command, giving it your module path -- here we use example.com/greetings. if you publish 
a module, this must be a path from which your module code can be downloaded by Go tools. That would be your 
code's repository

