package next_greatest_letter

import (
	"fmt"
	"testing"
)

type Case struct {
	Letters []byte
	Target  byte
	Output  byte
}

var cases = []Case{
	{
		[]byte{'c', 'f', 'j'},
		'a',
		'c',
	}, {
		[]byte{'c', 'f', 'j'},
		'c',
		'f',
	}, {
		[]byte{'c', 'f', 'j'},
		'd',
		'f',
	}, {
		[]byte{'c', 'f', 'j'},
		'g',
		'j',
	}, {
		[]byte{'c', 'f', 'j'},
		'j',
		'c',
	}, {
		[]byte{'c', 'f', 'j'},
		'k',
		'c',
	},
}

func TestNextGreatestLetter(t *testing.T) {
	for _, cs := range cases {
		if cs.Output != nextGreatestLetter(cs.Letters, cs.Target) {
			t.Error(fmt.Sprintf("%c error", cs.Target))
		}
	}
}
