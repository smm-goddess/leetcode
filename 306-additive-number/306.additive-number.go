package additive

/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 *
 * https://leetcode.com/problems/additive-number/description/
 *
 * algorithms
 * Medium (28.20%)
 * Total Accepted:    40.2K
 * Total Submissions: 142.2K
 * Testcase Example:  '"112358"'
 *
 * Additive number is a string whose digits can form additive sequence.
 *
 * A valid additive sequence should contain at least three numbers. Except for
 * the first two numbers, each subsequent number in the sequence must be the
 * sum of the preceding two.
 *
 * Given a string containing only digits '0'-'9', write a function to determine
 * if it's an additive number.
 *
 * Note: Numbers in the additive sequence cannot have leading zeros, so
 * sequence 1, 2, 03 or 1, 02, 3 is invalid.
 *
 * Example 1:
 *
 *
 * Input: "112358"
 * Output: true
 * Explanation: The digits can form an additive sequence: 1, 1, 2, 3, 5,
 * 8.
 * 1 + 1 = 2, 1 + 2 = 3, 2 + 3 = 5, 3 + 5 = 8
 *
 *
 * Example 2:
 *
 *
 * Input: "199100199"
 * Output: true
 * Explanation: The additive sequence is: 1, 99, 100, 199. 
 * 1 + 99 = 100, 99 + 100 = 199
 *
 * Follow up:
 * How would you handle overflow for very large input integers?
 */
func isAdditiveNumber(num string) bool {
	firstMaxLength, secondMaxLength, bs, checkedLength := len(num)/2, 0, []byte(num), 0
	for length := 1; length <= firstMaxLength; length++ {
		first := subString(bs, 0, length)
		secondMaxLength = (len(num) - length) / 2
		for sl := 1; sl <= secondMaxLength; sl++ {
			second := subString(bs, length, sl)
			if len(second) > 1 && second[0] == '0' {
				break
			}
			checkedLength = length + sl
			adder := append([]byte{}, first...)
			for {
				sum := byteSum(adder, second)
				if checkedLength+len(sum) <= len(bs) && equals(bs, sum, checkedLength) {
					checkedLength += len(sum)
					adder, second = second, sum
					if checkedLength == len(bs) {
						return true
					}
				} else {
					break
				}
			}
		}
	}
	return false
}

func equals(bs, target []byte, checkedLength int) bool {
	for i := 0; i < len(target); i++ {
		if bs[checkedLength+i] != target[i] {
			return false
		}
	}
	return true
}

func subString(bs []byte, start, length int) []byte {
	return bs[start : start+length]
}

func byteSum(a, b []byte) []byte {
	if len(a) < len(b) {
		a, b = b, a
	}
	var sum []byte
	var addFlag byte
	var i int
	for i = 1; i <= len(b); i++ {
		s := a[len(a)-i] - '0' + b[len(b)-i] - '0' + addFlag
		if s >= 10 {
			addFlag = 1
			sum = append([]byte{s - 10 + '0'}, sum...)
		} else {
			addFlag = 0
			sum = append([]byte{s + '0'}, sum...)
		}
	}
	for ; i <= len(a); i++ {
		s := a[len(a)-i] - '0' + addFlag
		if s >= 10 {
			addFlag = 1
			sum = append([]byte{s - 10 + '0'}, sum...)
		} else {
			addFlag = 0
			sum = append([]byte{s + '0'}, sum...)
		}
	}
	if addFlag == 1 {
		sum = append([]byte{'1'}, sum...)
	}
	return sum
}
