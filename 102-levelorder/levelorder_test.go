package levelorder

import "testing"

func TestLevelOrder(t *testing.T) {
	tree := &TreeNode{
		3,
		&TreeNode{9, nil, nil,},
		&TreeNode{
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	supposed := [][]int{{3}, {9, 20}, {15, 7}}
	level := levelOrder(tree)
	for i, sli := range level {
		for j, in := range sli {
			if supposed[i][j] != in {
				t.Error("Error")
				break
			}
		}
	}
}
