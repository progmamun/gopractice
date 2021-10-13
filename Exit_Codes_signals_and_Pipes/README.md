# Exit codes in bash

Every time a command gets executed in a shell, the resulting exit code gets
stored in a variable. The status of the last command executed gets stored in
the $? variable, and it can be printed as follows:
`> echo $? # will print 1`
It is important to note that the exit code only works when you run a binary
obtained with go build or go install. If you use go run, it will return 1 for any
code that is not 0.
