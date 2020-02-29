package goversion_test

import (
	"reflect"
	"testing"

	"github.com/sg0hsmt/goversion"
)

func TestVersionCompare(t *testing.T) {
	tbl := []struct {
		name string
		inA  *goversion.Version
		inB  *goversion.Version
		eq   bool // Equal()
	}{
		{
			name: "nil and nil",
			inA:  nil,
			inB:  nil,
			eq:   true,
		},
		{
			name: "nil and zero",
			inA:  nil,
			inB:  &goversion.Version{},
			eq:   false,
		},
		{
			name: "zero and nil",
			inA:  &goversion.Version{},
			inB:  nil,
			eq:   false,
		},
		{
			name: "zero and zero",
			inA:  &goversion.Version{},
			inB:  &goversion.Version{},
			eq:   true,
		},
		{
			name: "equal",
			inA:  &goversion.Version{Major: 1, Minor: 2, Patch: 3, Pre: ""},
			inB:  &goversion.Version{Major: 1, Minor: 2, Patch: 3, Pre: ""},
			eq:   true,
		},
		{
			name: "different major version",
			inA:  &goversion.Version{Major: 1, Minor: 2, Patch: 3, Pre: ""},
			inB:  &goversion.Version{Major: 0, Minor: 2, Patch: 3, Pre: ""},
			eq:   false,
		},
		{
			name: "different minor version",
			inA:  &goversion.Version{Major: 1, Minor: 2, Patch: 3, Pre: ""},
			inB:  &goversion.Version{Major: 1, Minor: 0, Patch: 3, Pre: ""},
			eq:   false,
		},
		{
			name: "different patch version",
			inA:  &goversion.Version{Major: 1, Minor: 2, Patch: 3, Pre: ""},
			inB:  &goversion.Version{Major: 1, Minor: 2, Patch: 0, Pre: ""},
			eq:   false,
		},
		{
			name: "different pre release",
			inA:  &goversion.Version{Major: 1, Minor: 2, Patch: 0, Pre: ""},
			inB:  &goversion.Version{Major: 1, Minor: 2, Patch: 0, Pre: "rc1"},
			eq:   false,
		},
	}

	for _, v := range tbl {
		t.Run(v.name, func(t *testing.T) {
			eq := v.inA.Equal(v.inB)
			if eq != v.eq {
				t.Errorf("unmatch equal, want %v, got %v", v.eq, eq)
			}
		})
	}
}

func TestReleaseTags(t *testing.T) {
	tbl := []struct {
		name string
		in   *goversion.Version
		out  []string
	}{
		{
			name: "empty",
			in:   &goversion.Version{},
			out:  []string{},
		},
		{
			name: "go1.13.8",
			in:   &goversion.Version{Major: 1, Minor: 13, Patch: 8, Pre: ""},
			out: []string{
				"go1.1",
				"go1.2",
				"go1.3",
				"go1.4",
				"go1.5",
				"go1.6",
				"go1.7",
				"go1.8",
				"go1.9",
				"go1.10",
				"go1.11",
				"go1.12",
				"go1.13",
			},
		},
		{
			name: "go1.14beta1",
			in:   &goversion.Version{Major: 1, Minor: 14, Patch: 0, Pre: "beta1"},
			out: []string{
				"go1.1",
				"go1.2",
				"go1.3",
				"go1.4",
				"go1.5",
				"go1.6",
				"go1.7",
				"go1.8",
				"go1.9",
				"go1.10",
				"go1.11",
				"go1.12",
				"go1.13",
				"go1.14",
			},
		},
		{
			name: "empty",
			in:   &goversion.Version{},
			out:  []string{},
		},
	}

	for _, v := range tbl {
		t.Run(v.name, func(t *testing.T) {
			res := v.in.ReleaseTags()
			if !reflect.DeepEqual(res, v.out) {
				t.Errorf("unmatch release tags, want %v, got %v", v.out, res)
			}
		})
	}
}
