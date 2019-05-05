package findrepeateddnasequences

import (
	"fmt"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	sequences := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	fmt.Println(sequences)
	if sequences[0] != "AAAAACCCCC" || sequences[1] != "CCCCCAAAAA" {
		t.Error("Error")
	}
}
