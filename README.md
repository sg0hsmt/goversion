# goversion

[![GoDoc](https://godoc.org/github.com/sg0hsmt/goversion?status.svg)](https://godoc.org/github.com/sg0hsmt/goversion)
[![Test](https://github.com/sg0hsmt/goversion/workflows/Test/badge.svg)](https://github.com/sg0hsmt/goversion/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/sg0hsmt/goversion)](https://goreportcard.com/report/github.com/sg0hsmt/goversion)
[![codecov](https://codecov.io/gh/sg0hsmt/goversion/branch/master/graph/badge.svg)](https://codecov.io/gh/sg0hsmt/goversion)
[![License](https://img.shields.io/github/license/sg0hsmt/goversion.svg)](https://github.com/sg0hsmt/goversion/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/sg0hsmt/goversion.svg)](https://github.com/sg0hsmt/goversion/releases/latest)

`goversion` is a package that gets the go version from the go command.

## How to use

See godoc examples.

```go
package main

import (
	"fmt"

	"github.com/sg0hsmt/goversion"
)

func main() {
	ver, err := goversion.Discover()
	if err != nil {
		fmt.Printf("discover failed: %v", err)
		return
	}

	fmt.Printf("version: %s\n", ver)
	fmt.Printf("major: %d\n", ver.Major)
	fmt.Printf("minor: %d\n", ver.Minor)
	fmt.Printf("patch: %d\n", ver.Patch)

	if ver.IsPreRelease() {
		fmt.Printf("pre: %q\n", ver.Pre)
	}
}
```

## Limitations

The development version (called gotip) is not supported.
