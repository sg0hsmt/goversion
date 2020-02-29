package goversion

import (
	"fmt"
	"strings"
)

// Version represent the go version information.
type Version struct {
	Major int    `json:"major"`
	Minor int    `json:"minor"`
	Patch int    `json:"patch,omitempty"`
	Pre   string `json:"prerelease,omitempty"`
}

// Clone returns a deep copy of instance.
func (v *Version) Clone() *Version {
	return &Version{
		Major: v.Major,
		Minor: v.Minor,
		Patch: v.Patch,
		Pre:   v.Pre,
	}
}

// String returns go version string.
func (v *Version) String() string {
	if v.IsPreRelease() {
		return fmt.Sprintf("go%d.%d%s", v.Major, v.Minor, v.Pre)
	}

	if v.Patch == 0 {
		return fmt.Sprintf("go%d.%d", v.Major, v.Minor)
	}

	return fmt.Sprintf("go%d.%d.%d", v.Major, v.Minor, v.Patch)
}

// BuildTag returns go build tag string.
func (v *Version) BuildTag() string {
	return fmt.Sprintf("go%d.%d", v.Major, v.Minor)
}

// ReleaseTags returns release build tags.
// see https://golang.org/pkg/go/build/#Context
func (v *Version) ReleaseTags() []string {
	res := make([]string, 0, v.Minor)

	for i := 1; i <= v.Minor; i++ {
		res = append(res, fmt.Sprintf("go1.%d", i))
	}

	return res
}

// IsPreRelease reports whether pre release version.
func (v *Version) IsPreRelease() bool {
	return v.Pre != ""
}

// IsBeta reports whether beta version.
func (v *Version) IsBeta() bool {
	return strings.HasPrefix(v.Pre, "beta")
}

// IsRC reports whether release candidate version.
func (v *Version) IsRC() bool {
	return strings.HasPrefix(v.Pre, "rc")
}

// Equal reports whether v and u represent the same version instant.
func (v *Version) Equal(u *Version) bool {
	if v == nil && u == nil {
		return true
	}

	if v == nil || u == nil {
		return false
	}

	return v.Major == u.Major && v.Minor == u.Minor && v.Patch == u.Patch && v.Pre == u.Pre
}
