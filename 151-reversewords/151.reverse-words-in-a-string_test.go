package reversewords

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	if reverseWords("the sky is blue") != "blue is sky the" {
		t.Error("Sky blue Error")
	}
	if reverseWords("  hello world!  ") != "world! hello" {
		t.Error("Hello World Error")
	}
	if reverseWords("a good   example") != "example good a" {
		t.Error("Example Error")
	}
	if reverseWords("") != "" {
		t.Error("Empty Error")
	}
	if reverseWords(" ") != "" {
		t.Error("Space Error")
	}
}

//func TestSplit(t *testing.T) {
//	fmt.Println(split("the sky is blue"))
//	fmt.Println(split("  hello world!  "))
//	fmt.Println(split("a good   example"))
//}
