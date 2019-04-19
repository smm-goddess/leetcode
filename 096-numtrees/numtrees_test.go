package numtrees

import (
	"fmt"
	"testing"
)

var ns = map[int]int{1: 1, 2: 2, 3: 5, 4: 14, 5: 42}

func TestNumTrees(t *testing.T) {
	for k, v := range ns {
		if numTrees(k) != v {
			t.Error(fmt.Sprintf("Test Fail %d", k))
		}
	}
}
