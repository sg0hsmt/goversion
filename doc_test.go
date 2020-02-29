package goversion_test

import (
	"fmt"

	"github.com/sg0hsmt/goversion"
)

func Example() {
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
