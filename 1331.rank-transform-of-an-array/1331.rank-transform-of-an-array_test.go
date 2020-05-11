package arrayRankTransform

import (
	"testing"
)

func arrayEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}else {
		for index, i := range a {
			if b[index] != i {
				return false
			}
		}
		return true
	}
}

func TestArrayRankTransform(t *testing.T){
	// if !arrayEqual(arrayRankTransform([]int{40,10,20,30}), []int{4,1,2,3}){
	// 	t.Error("error")
	// }
	// if !arrayEqual(arrayRankTransform([]int{100,100,100}), []int{1,1,1}){
	// 	t.Error("error 2")
	// }
	// if !arrayEqual(arrayRankTransform([]int{37,12,28,9,100,56,80,5,12}), []int{5,3,4,2,8,6,7,1,3}){
	// 	t.Error("error 3")
	// }
	// if !arrayEqual(arrayRankTransform([]int{37,12,28,9,100,56,56,80,5,12,5}), []int{5,3,4,2,8,6,6,7,1,3,1}){
	// 	t.Error("error 4")
	// }
	// if !arrayEqual(arrayRankTransform([]int{14,-19,12,-25,34,-27,-48,-37,14,-47,40,23,46,-39,48,-41,18,-27,-4}), 
	// 								  []int{11, 8, 10,  7,14,  6,  1,  5,11,  2,15,13,16,  4,17,  3,12,  6, 9}){
	// 								//	    11  8  10   7 14   6   1   5 11   2 16 13 17   4 18   3 12   6  9
	// 	t.Error("error 4")
	// }
	if !arrayEqual(arrayRankTransform([]int{-34,-47,40,3 ,-30,29,-16,-45,2,25,48,23,2,-43, 4,-25,26,-19,44, 3,-30,-43}), 
									  []int{  4,  1,16,10,  5,15,  8,  2,9,13,18,12,9,  3,11,  6,14,  7,17,10,  5,  3}){
									//	      4,  1,16,10,  5,15,  8,  2,9,13,18,12,9,  3,11,  6,14,  7,16,10,  5,  3]
		t.Error("error 4")
	}
}
