# Call your code from another module

Here I tried to create another module to call a function that we created in the create-module branch
In order to use that func. Created a new folder named hello. in there a new go mod initialized,
after that process, we created a new go file. to import previously written code. As you can see in the doc
we imported a module by the name we use to initialize it. in my case it was example.com/greetings.
In order go mod package manager to install that module, we have to publish it on go tools, after that
it will be available for usage. For now, as we didn't publish the module to the go tools, we have to
somehow tell package manager, if you see example.com/greetings go above the folder. We can do this by
this command

```bash
$ go mod edit -replace example.com/greetings=../
```

The command specifies that example.com/greetings should be replaced with ../ for the purpose of locating
dependency.

We use this command to download to synchronize the module dependencies, adding those required by the code,
but not yet tracked in the module.

```bash
 $ go mod tidy 
```

After we run the command go mod file should look like this:

```mod
module example.com/hello

go 1.22.1

replace example.com/greetings => ../

require example.com/greetings v0.0.0-00010101000000-000000000000

```

The command found the local code in the upper directory, the added a require directive to specify that example.com/hello
requires example.com/greetings. We created this when we imported the greetings package in hello.go

The number following the module path is a pseudo-version number -- a generated number used in place of a semantic
version number (which the module doesn't have yet).

To reference a published module, a go.mod file would typically omit the replace directive and use a require directive
with a tagged version number at the end.

```mod
require example.com/greetings v1.1.0
```