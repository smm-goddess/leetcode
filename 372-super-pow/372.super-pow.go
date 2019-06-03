package superpow

/*
 * @lc app=leetcode id=372 lang=golang
 *
 * [372] Super Pow
 *
 * https://leetcode.com/problems/super-pow/description/
 *
 * algorithms
 * Medium (35.47%)
 * Total Accepted:    27.6K
 * Total Submissions: 77.5K
 * Testcase Example:  '2\n[3]'
 *
 * Your task is to calculate ab mod 1337 where a is a positive integer and b is
 * an extremely large positive integer given in the form of an array.
 *
 * Example 1:
 *
 *
 *
 * Input: a = 2, b = [3]
 * Output: 8
 *
 *
 *
 * Example 2:
 *
 *
 * Input: a = 2, b = [1,0]
 * Output: 1024
 *
 *
 *
 */
func superPow(a int, b []int) int {
	if b == nil || len(b) == 0 || a == 1 {
		return 1
	} else if a == 0 {
		return 0
	}
	return superPowIter(1, a, b)
}

func superPowIter(iter, a int, b []int) int {
	if len(b) == 0 {
		return iter % 1337
	} else {
		if b[len(b)-1]&1 == 0 {
			b = half(b)
			return superPowIter(iter, (a*a)%1337, b)
		} else {
			b[len(b)-1]--
			return superPowIter((iter*a)%1337, a, b)
		}
	}
}

func half(b []int) []int {
	var ans []int
	var remains int
	if b[0] >= 2 {
		remains = b[0] % 2
		ans = append(ans, b[0]/2)
	} else {
		remains = b[0]
	}
	for i := 1; i < len(b); i++ {
		ans = append(ans, (remains*10+b[i])/2)
		remains = (remains*10 + b[i]) % 2
	}
	return ans
}
