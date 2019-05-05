package rightsideview

import "testing"

func intSliceEquals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, c := range a {
			if b[i] != c {
				return false
			}
		}
		return true
	}
}

func TestRightSideView(t *testing.T) {
	tree := &TreeNode{1,
		&TreeNode{2,
			nil,
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			nil,
			&TreeNode{4, nil, nil}}}
	if !intSliceEquals(rightSideView(tree), []int{1, 3, 4}) {
		t.Error("1 3 4 Error")
	}
}
