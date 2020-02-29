package goversion

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// ErrDevelopVersion develop version is not supported.
var ErrDevelopVersion = fmt.Errorf("develop version is not supported")

// ErrVersionSyntax version syntax parse failed.
var ErrVersionSyntax = fmt.Errorf("version syntax parse failed")

// nolint: gochecknoglobals
var mu sync.Mutex

// nolint: gochecknoglobals
var cache *Version

// Discover returns version instance from go command execute result.
// Execution results are cached and reused.
func Discover() (*Version, error) {
	mu.Lock()
	defer mu.Unlock()

	if cache == nil {
		v, err := discover()
		if err != nil {
			return nil, err
		}

		cache = v
	}

	return cache.Clone(), nil
}

// discover returns version instance from go command execute result.
func discover() (*Version, error) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		return nil, err
	}

	return parse(string(out))
}

// parse returns version instance from go version output.
func parse(s string) (*Version, error) {
	if strings.HasPrefix(s, "go version devel ") {
		return nil, ErrDevelopVersion
	}

	re := regexp.MustCompile(`go version go(\d+)\.(\d+)(\.|[a-z]+[0-9]+|)(\d+|) .+/.+`)

	v := re.FindStringSubmatch(s)
	if v == nil {
		return nil, ErrVersionSyntax
	}

	major, _ := strconv.Atoi(v[1])
	minor, _ := strconv.Atoi(v[2])

	var patch int
	if v[3] == "." && v[4] != "" {
		patch, _ = strconv.Atoi(v[4])
	}

	var pre string
	if v[3] != "." {
		pre = v[3]
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Pre:   pre,
	}, nil
}
