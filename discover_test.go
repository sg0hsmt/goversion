package goversion_test

import (
	"errors"
	"go/build"
	"reflect"
	"runtime"
	"testing"

	"github.com/sg0hsmt/goversion"
)

func TestDiscover(t *testing.T) {
	var first *goversion.Version

	t.Run("first call (get from go command)", func(t *testing.T) {
		res, err := goversion.Discover()
		if err != nil {
			t.Fatalf("discover is failed, %v", err)
		}

		runver := runtime.Version()
		if res.String() != runver {
			t.Errorf("unmatch version string, discover %q, runtime %q", res.String(), runver)
		}

		reltags := build.Default.ReleaseTags
		if !reflect.DeepEqual(res.ReleaseTags(), reltags) {
			t.Errorf("unmatch release tags, discover %q, runtime %q", res.ReleaseTags(), reltags)
		}

		first = res
	})

	var testErr = errors.New("test for execution failure")

	// replace internal function
	goversion.SetExecCmd(func() ([]byte, error) {
		return nil, testErr // always error
	})

	t.Run("second call (get from cache)", func(t *testing.T) {
		res, err := goversion.Discover()
		if err != nil {
			t.Fatalf("discover is failed, %v", err)
		}

		if !first.Equal(res) {
			t.Errorf("second is not equal first, first %#v, second %#v", first, res)
		}
	})

	// clear internal cache
	goversion.ResetCache()

	t.Run("third call (from go command)", func(t *testing.T) {
		_, err := goversion.Discover()
		if err != testErr {
			t.Errorf("unmatch error, want %v, got %v", testErr, err)
		}
	})
}
