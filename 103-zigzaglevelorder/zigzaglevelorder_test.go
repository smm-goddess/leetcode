package zigzaglevelorder

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
	supposed := [][]int{{3}, {20, 9}, {15, 7}}
	level := zigzagLevelOrder(tree)
	for i, sli := range level {
		for j, in := range sli {
			if supposed[i][j] != in {
				t.Error("Error")
				break
			}
		}
	}
	tree = &TreeNode{
		1,
		&TreeNode{
			2,
			&TreeNode{5, nil, nil},
			&TreeNode{11, nil, nil},
		},
		&TreeNode{3,
			&TreeNode{9,
				&TreeNode{6, nil, nil},
				&TreeNode{8, nil, nil}},
			&TreeNode{20,
				&TreeNode{15, nil, nil},
				&TreeNode{7, nil, nil},
			},
		},
	}
	supposed = [][]int{{1}, {3, 2}, {5, 11, 9, 20}, {7, 15, 8, 6}}
	level = zigzagLevelOrder(tree)
	for i, sli := range level {
		for j, in := range sli {
			if supposed[i][j] != in {
				t.Error("Error")
				break
			}
		}
	}
}
