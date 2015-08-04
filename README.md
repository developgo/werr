# Error Wrapper

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/txgruppi/werr)
![Codeship](https://img.shields.io/codeship/563848f0-1cc7-0133-3afb-0ee7adf4cd2d.svg?style=flat-square)
[![Codecov](https://img.shields.io/codecov/c/github/txgruppi/werr.svg?style=flat-square)](https://codecov.io/github/txgruppi/werr)

Error Wrapper creates an wrapper for the `error` type in Go which captures the File, Line and Stack of where it was called.

## Why?

I don't like to use `panic`.

I want my apps to run **forever** and just output errors to a log.

Usually I write my logs using `logger.Printf("%#v", err)` but I needed more info related to the error, so I created this package.

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/txgruppi/werr"
)

func main() {
	file, err := os.Open("/tmp/my-missing-file")
	if err != nil {
		err = werr.Wrap(err)                        // Wrap it
		fmt.Println(err.Error())                    // Return the original error message
		if wrapped, ok := err.(*werr.Wrapper); ok { // Try to convert to `*werr.Wrapper`
			lg, _ := wrapped.Log() // Generate the log message
			fmt.Println(lg)        // Print the log message
		}
	}
	defer file.Close()
}
```

This code will output something line this:

```
open /tmp/my-missing-file: no such file or directory
/Users/txgruppi/code/temp/test.go:13 open /tmp/my-missing-file: no such file or directory
goroutine 1 [running]:
github.com/txgruppi/werr.Wrap(0x2208246900, 0x208270300, 0x0, 0x0)
        /Users/txgruppi/code/go/src/github.com/txgruppi/werr/funcs.go:24 +0x153
main.main()
        /Users/txgruppi/code/temp/test.go:13 +0x72
```
