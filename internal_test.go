package goversion

import "testing"

func TestParse(t *testing.T) {
	tbl := []struct {
		name string
		in   string
		out  *Version
		err  error

		str  string // String()
		tag  string // BuildTag()
		pre  bool   // IsPreRelease()
		beta bool   // IsBeta()
		rc   bool   // IsRC()
	}{
		{
			name: "empty",
			in:   "",
			out:  nil,
			err:  ErrVersionSyntax,
		},
		{
			name: "one digit",
			in:   "go version go1.2.3 goos/goarch",
			out:  &Version{Major: 1, Minor: 2, Patch: 3},
			err:  nil,
			str:  "go1.2.3",
			tag:  "go1.2",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			name: "three digit",
			in:   "go version go123.456.789 goos/goarch",
			out:  &Version{Major: 123, Minor: 456, Patch: 789},
			err:  nil,
			str:  "go123.456.789",
			tag:  "go123.456",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			name: "major is not numerical",
			in:   "go version goX.2.3 goos/goarch",
			out:  nil,
			err:  ErrVersionSyntax,
		},
		{
			name: "minor is not numerical",
			in:   "go version go1.Y.3 goos/goarch",
			out:  nil,
			err:  ErrVersionSyntax,
		},
		{
			name: "patch is not numerical",
			in:   "go version go1.2.Z goos/goarch",
			out:  nil,
			err:  ErrVersionSyntax,
		},
		{
			name: "no patch version",
			in:   "go version go1.2 goos/goarch",
			out:  &Version{Major: 1, Minor: 2, Patch: 0},
			err:  nil,
			str:  "go1.2",
			tag:  "go1.2",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			name: "pre release",
			in:   "go version go1.2alpha0 goos/goarch",
			out:  &Version{Major: 1, Minor: 2, Patch: 0, Pre: "alpha0"},
			err:  nil,
			str:  "go1.2alpha0",
			tag:  "go1.2",
			pre:  true,
			beta: false,
			rc:   false,
		},
		{
			name: "go 1.13",
			in:   "go version go1.13 goos/goarch",
			out:  &Version{Major: 1, Minor: 13, Patch: 0},
			err:  nil,
			str:  "go1.13",
			tag:  "go1.13",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			name: "go 1.13.8",
			in:   "go version go1.13.8 goos/goarch",
			out:  &Version{Major: 1, Minor: 13, Patch: 8},
			err:  nil,
			str:  "go1.13.8",
			tag:  "go1.13",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			name: "go1.14beta1",
			in:   "go version go1.14beta1 goos/goarch",
			out:  &Version{Major: 1, Minor: 14, Patch: 0, Pre: "beta1"},
			err:  nil,
			str:  "go1.14beta1",
			tag:  "go1.14",
			pre:  true,
			beta: true,
			rc:   false,
		},
		{
			name: "go1.14rc1",
			in:   "go version go1.14rc1 goos/goarch",
			out:  &Version{Major: 1, Minor: 14, Patch: 0, Pre: "rc1"},
			err:  nil,
			str:  "go1.14rc1",
			tag:  "go1.14",
			pre:  true,
			beta: false,
			rc:   true,
		},
		{
			name: "go1.14",
			in:   "go version go1.14 goos/goarch",
			out:  &Version{Major: 1, Minor: 14, Patch: 0},
			err:  nil,
			str:  "go1.14",
			tag:  "go1.14",
			pre:  false,
			beta: false,
			rc:   false,
		},
		{
			// not supported.
			name: "gotip",
			in:   "go version devel +5756808 Sat Feb 29 10:21:33 2020 +0000 goos/goarch",
			out:  nil,
			err:  ErrDevelopVersion,
		},
	}

	for _, v := range tbl {
		t.Run(v.name, func(t *testing.T) {
			res, err := parse(v.in)

			if err == nil && v.err != nil {
				t.Errorf("parse is success, want %v", v.err)

				return
			}

			if err != nil && v.err == nil {
				t.Errorf("parse is failed, got %v", err)

				return
			}

			if err != nil && v.err != nil {
				if res != nil {
					t.Errorf("parse is failed, but version is not nil")
				}
				if err != v.err {
					t.Errorf("unmatch error, want %v, got %v", v.err, err)
				}

				return
			}

			if !res.Equal(v.out) {
				t.Errorf("unmatch result, want %v, got %v", v.out, res)
			}

			if res.String() != v.str {
				t.Errorf("unmatch string, want %q, got %q", v.str, res.String())
			}

			if res.BuildTag() != v.tag {
				t.Errorf("unmatch build tag, want %q, got %q", v.tag, res.BuildTag())
			}

			if res.IsPreRelease() != v.pre {
				t.Errorf("unmatch pre release, want %v, got %v", v.pre, res.IsPreRelease())
			}

			if res.IsBeta() != v.beta {
				t.Errorf("unmatch beta, want %v, got %v", v.pre, res.IsBeta())
			}

			if res.IsRC() != v.rc {
				t.Errorf("unmatch release candidate, want %v, got %v", v.pre, res.IsRC())
			}
		})
	}
}
