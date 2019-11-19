package license_key_formatting

import (
	"testing"
)

type TestCase struct {
	s      string
	k      int
	output string
}

var testCases []*TestCase

func init() {
	testCases = []*TestCase{
		{
			"5F3Z-2e-9-w",
			4,
			"5F3Z-2E9W",
		},
		{
			"2-5g-3-J",
			2,
			"2-5G-3J",
		},
		{
			"2-5g-3-J",
			1,
			"2-5-G-3-J",
		},
		{
			"2-5g-3-J",
			6,
			"25G3J",
		},
		{
			"a-a-a-a-",
			1,
			"A-A-A-A",
		},
		{
			"---",
			3,
			"",
		},
	}
}

func TestLicenseFormatting(t *testing.T) {
	for _, cs := range testCases {
		if cs.output != licenseKeyFormatting(cs.s, cs.k) {
			t.Error(cs.s, " ", cs.k, " error")
		}
	}
}
