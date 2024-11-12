Start off bash scripts with the following shebang:
#!/bin/bash

Declaring variables - Note that there cannot be any spaces between
the variable name, the equals sign, and the value
```
variable="text"
```

String operations
```
echo "${variable:i:j}" # returns characters from [i,j)
echo "${variable: -i}" # returns last i characters
echo "${#variable}"    # returns length of string
```

Arrays
```
array=(this is an array)
echo "${array[i]}" # print element at index i
echo "${array[@]}" # prints all elements
```
