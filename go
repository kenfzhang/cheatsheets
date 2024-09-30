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
