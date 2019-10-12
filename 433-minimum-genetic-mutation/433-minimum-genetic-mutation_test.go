package minimum_genetic_mutation

import (
	"fmt"
	"testing"
)

func TestMinMutation(t *testing.T) {
	if minMutation("AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}) != 1 {
		t.Error("AACCGGTT error")
	}
	if minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}) != 2 {
		t.Error("AACCGGTT error")
	}
	if minMutation("AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}) != 3 {
		t.Error("AAAAACCC error")
	}
	if distance := minMutation("AACCTTGG", "AATTCCGG", []string{"AATTCCGG", "AACCTGGG", "AACCCCGG", "AACCTACC"}); distance != -1 {
		t.Error(fmt.Sprintf("AACCTTGG error %d", distance))
	}
	if minMutation("AACCGGTT", "AACCGCTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}) != 2 {
		t.Error("error")
	}
}
