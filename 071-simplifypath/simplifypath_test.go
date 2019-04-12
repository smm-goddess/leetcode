package simplifypath

import (
	"fmt"
	"testing"
)

var testMap = map[string]string{"/../": "/",
	"/home//foo/":           "/home/foo",
	"/home/":                "/home",
	"/a/./b/../../c/":       "/c",
	"/a/../../b/../c//.//":  "/c",
	"/a//b////c/d//././/..": "/a/b/c"}

func TestSimplifyPath(t *testing.T) {
	for k, v := range testMap {
		if simplifyPath(k) != v {
			t.Error(fmt.Sprintf("%s Error", k))
		}
	}
}
