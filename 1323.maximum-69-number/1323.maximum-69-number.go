
package maximun_69_number

import "strconv"

func maximum69Number (num int) int {
	ans, found6 := 0, false
	s :=  strconv.Itoa(num)
	for _, c := range []byte(s){
		if !found6 && c == '6' {
			ans = ans * 10 + 9
			found6 = true
		} else {
			ans = ans * 10 + int(c - '0')
		}
	}
	return ans
}