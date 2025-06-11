package fail

fail

/*
This is a non-compiling file that has been added to explicitly ensure that CI fails.
It also contains the command that caused the failure and its output.
Remove this file if debugging locally.

go mod operation failed. This may mean that there are legitimate dependency issues with the "go.mod" definition in the repository and the updates performed by the gomod check. This branch can be cloned locally to debug the issue.

Command that caused error:
./godelw check compiles

Output:
Running compiles...
nobadfuncs/creator/creator.go:34:11: cannot use cfg.ToChecker() (value of type *nobadfuncs.Checker) as okgo.Checker value in return statement: *nobadfuncs.Checker does not implement okgo.Checker (missing method MultiCPU)
Finished compiles
Check(s) produced output: [compiles]

*/
