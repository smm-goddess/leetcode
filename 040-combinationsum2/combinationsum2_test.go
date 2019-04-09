package combinationsum2

import "testing"

func TestCombinationSum(t *testing.T){
  candidate, target := []int{10, 1, 2, 7, 6, 1, 5}, 8 
  t.Log(combinationSum2(candidate, target))
  candidate, target = []int{2, 5, 2, 1, 2}, 5
  t.Log(combinationSum2(candidate, target))
}
