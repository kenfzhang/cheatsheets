Run .zig file containing main()
$ zig run <filename>.zig

Print
```
std.debug.print("string {variable}", .{"variable"});
```

Running tests
$ zig test <filename>.zig

Create a test
```
const std = @import("std");
const expect = std.testing.expect;

test "always succeeds" {
	try expect(true);
}

test "always fails" {
	try expect(false);
}
```
