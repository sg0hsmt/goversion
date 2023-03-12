package goversion_test

import (
	"fmt"

	"github.com/sg0hsmt/goversion"
)

//nolint:testableexamples
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

func ExampleVersion_Clone() {
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}

	clone := go113.Clone()

	fmt.Printf("go113 == clone = %v\n", go113 == clone)
	fmt.Printf("go113.Equal(clone) = %v\n", go113.Equal(clone))
	fmt.Printf("go113: %#v\n", go113)
	fmt.Printf("clone: %#v\n", clone)

	// Output:
	// go113 == clone = false
	// go113.Equal(clone) = true
	// go113: &goversion.Version{Major:1, Minor:13, Patch:8, Pre:""}
	// clone: &goversion.Version{Major:1, Minor:13, Patch:8, Pre:""}
}

func ExampleVersion_String() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.String() = %q\n", go113.String())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.String() = %q\n", go114beta1.String())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.String() = %q\n", go114rc1.String())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.String() = %q\n", go114.String())

	// Output:
	// go113.String() = "go1.13.8"
	// go114beta1.String() = "go1.14beta1"
	// go114rc1.String() = "go1.14rc1"
	// go114.String() = "go1.14"
}

func ExampleVersion_BuildTag() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.BuildTag() = %q\n", go113.BuildTag())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.BuildTag() = %q\n", go114beta1.BuildTag())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.BuildTag() = %q\n", go114rc1.BuildTag())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.BuildTag() = %q\n", go114.BuildTag())

	// Output:
	// go113.BuildTag() = "go1.13"
	// go114beta1.BuildTag() = "go1.14"
	// go114rc1.BuildTag() = "go1.14"
	// go114.BuildTag() = "go1.14"
}

//nolint:lll
func ExampleVersion_ReleaseTags() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.ReleaseTags() = %v\n", go113.ReleaseTags())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.ReleaseTags() = %v\n", go114beta1.ReleaseTags())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.ReleaseTags() = %v\n", go114rc1.ReleaseTags())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.ReleaseTags() = %v\n", go114.ReleaseTags())

	// Output:
	// go113.ReleaseTags() = [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11 go1.12 go1.13]
	// go114beta1.ReleaseTags() = [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11 go1.12 go1.13 go1.14]
	// go114rc1.ReleaseTags() = [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11 go1.12 go1.13 go1.14]
	// go114.ReleaseTags() = [go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11 go1.12 go1.13 go1.14]
}

func ExampleVersion_IsPreRelease() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.IsPreRelease() = %v\n", go113.IsPreRelease())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.IsPreRelease() = %v\n", go114beta1.IsPreRelease())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.IsPreRelease() = %v\n", go114rc1.IsPreRelease())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.IsPreRelease() = %v\n", go114.IsPreRelease())

	// Output:
	// go113.IsPreRelease() = false
	// go114beta1.IsPreRelease() = true
	// go114rc1.IsPreRelease() = true
	// go114.IsPreRelease() = false
}

func ExampleVersion_IsBeta() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.IsBeta() = %v\n", go113.IsBeta())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.IsBeta() = %v\n", go114beta1.IsBeta())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.IsBeta() = %v\n", go114rc1.IsBeta())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.IsBeta() = %v\n", go114.IsBeta())

	// Output:
	// go113.IsBeta() = false
	// go114beta1.IsBeta() = true
	// go114rc1.IsBeta() = false
	// go114.IsBeta() = false
}

func ExampleVersion_IsRC() {
	// go 1.13.8
	go113 := &goversion.Version{
		Major: 1,
		Minor: 13,
		Patch: 8,
	}
	fmt.Printf("go113.IsRC() = %v\n", go113.IsRC())

	// go 1.14 beta1
	go114beta1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "beta1",
	}
	fmt.Printf("go114beta1.IsRC() = %v\n", go114beta1.IsRC())

	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}
	fmt.Printf("go114rc1.IsRC() = %v\n", go114rc1.IsRC())

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}
	fmt.Printf("go114.IsRC() = %v\n", go114.IsRC())

	// Output:
	// go113.IsRC() = false
	// go114beta1.IsRC() = false
	// go114rc1.IsRC() = true
	// go114.IsRC() = false
}

func ExampleVersion_Equal() {
	// go 1.14 rc1
	go114rc1 := &goversion.Version{
		Major: 1,
		Minor: 14,
		Pre:   "rc1",
	}

	// go 1.14
	go114 := &goversion.Version{
		Major: 1,
		Minor: 14,
	}

	clone := go114.Clone()

	fmt.Printf("go114.Equal(go1.14) = %v\n", go114.Equal(go114))
	fmt.Printf("go114.Equal(clone) = %v\n", go114.Equal(clone))
	fmt.Printf("go114.Equal(go1.14rc1) = %v\n", go114.Equal(go114rc1))

	// Output:
	// go114.Equal(go1.14) = true
	// go114.Equal(clone) = true
	// go114.Equal(go1.14rc1) = false
}
