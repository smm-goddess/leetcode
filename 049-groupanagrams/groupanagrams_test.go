package groupanagrams

import "testing"

func TestGroupAnagrams(t *testing.T) {
	origin := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	t.Log(groupAnagrams(origin))
}
