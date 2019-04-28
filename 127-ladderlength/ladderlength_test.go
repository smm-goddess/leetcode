package ladderlength

import (
	"fmt"
	"testing"
)

func TestLadderLength(t *testing.T) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	if ladderLength(beginWord, endWord, wordList) != 5 {
		t.Error("Fail 1")
	}

	wordList = wordList[:len(wordList)-1]
	if ladderLength(beginWord, endWord, wordList) != 0 {
		t.Error("Fail 2")
	}
	wordList = []string{}
	if ladderLength(beginWord, endWord, wordList) != 0 {
		t.Error("Fail 3")
	}
	beginWord = "hot"
	endWord = "dog"
	wordList = []string{"hot", "cog", "dog", "tot", "hog", "hop", "pot", "dot"}
	if ladderLength(beginWord, endWord, wordList) != 3 {
		t.Error("Fail 4")
	}
	beginWord = "qa"
	endWord = "sq"
	wordList = []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}
	if ladderLength(beginWord, endWord, wordList) != 5 {
		t.Error(fmt.Sprintf("Fail 5, %d", ladderLength(beginWord, endWord, wordList)))
	}
}
