Initialize new project (from root dir)
$ go mod init <prefix/title>

Run project (from root dir)
$ go run .

Build and create executable (from root dir)
$ go build <.go file>

Install a dependency
$ go get -u <dependency>
Where <dependency> is something like github.com/go-chi/chi/v5
The dependency should now be suggested by autocomplete

Note: it is preferred to use 'go install' instead of 'go get':
$ go install <dependency>

To load modules stored locally, add a line with the following
format to go.mod:
```
replace <import name> => <local import path>
```
Where <import name> is the name used in the import statement,
and <local import path> is the path of the module from the root
of the current module. For example:
```
replace example.com/greetings => ../greetings
```
The above line makes it so importing "example.com/greetings"
looks in the directory "../greetings" relative to the current directory.

Initialization statements in if statements allow you to declare and
initialize a variable only within the scope of that if-block.
Note: only one variable can be declared in the init statement.
```
if <inititializer>; <condition> {
	// execute if condition is true
}
```
