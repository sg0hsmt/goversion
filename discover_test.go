package goversion_test

import (
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

	t.Run("second call (get from cache)", func(t *testing.T) {
		res, err := goversion.Discover()
		if err != nil {
			t.Fatalf("discover is failed, %v", err)
		}

		if !first.Equal(res) {
			t.Errorf("second is not equal first, first %#v, second %#v", first, res)
		}
	})
}
