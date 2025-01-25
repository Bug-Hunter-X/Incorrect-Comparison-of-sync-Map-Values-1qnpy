# Incorrect Comparison of sync.Map Values in Go
This example demonstrates a common error when working with Go's `sync.Map`.  The `sync.Map.Load` method returns an `interface{}`, requiring a type assertion before comparison.

The `bug.go` file contains code that attempts to directly compare the loaded value with a string literal, resulting in a false negative.
The `bugSolution.go` file shows the corrected way to perform the comparison using type assertion.