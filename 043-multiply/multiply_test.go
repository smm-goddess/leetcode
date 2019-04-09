package multiply

import "testing"

func TestMultiply(t *testing.T){
  s1, s2 := “0”，“123”
  if multiply(s1, s2) == "0" {
    t.Log("Pass")
  }
}
