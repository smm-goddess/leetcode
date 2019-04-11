package inttoroman

import (
	"bytes"
)

/*
 * @lc app=leetcode id=12 lang=golang
 *
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (50.26%)
 * Total Accepted:    216.8K
 * Total Submissions: 430.4K
 * Testcase Example:  '3'
 *
 * Roman numerals are represented by seven different symbols: I, V, X, L, C, D
 * and M.
 *
 *
 * Symbol       Value
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * For example, two is written as II in Roman numeral, just two one's added
 * together. Twelve is written as, XII, which is simply X + II. The number
 * twenty seven is written as XXVII, which is XX + V + II.
 *
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII. Instead, the number four is
 * written as IV. Because the one is before the five we subtract it making
 * four. The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 *
 *
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 *
 *
 * Given an integer, convert it to a roman numeral. Input is guaranteed to be
 * within the range from 1 to 3999.
 *
 * Example 1:
 *
 *
 * Input: 3
 * Output: "III"
 *
 * Example 2:
 *
 *
 * Input: 4
 * Output: "IV"
 *
 * Example 3:
 *
 *
 * Input: 9
 * Output: "IX"
 *
 * Example 4:
 *
 *
 * Input: 58
 * Output: "LVIII"
 * Explanation: L = 50, V = 5, III = 3.
 *
 *
 * Example 5:
 *
 *
 * Input: 1994
 * Output: "MCMXCIV"
 * Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
 *
 */
func intToRoman(num int) string {
	divisor := 1000
	for num < divisor {
		divisor = divisor / 10
	}
	remain := num / divisor
	result, index := make([]byte, 16), 0
	for num > 0 && divisor > 0 {
		num = num - remain*divisor
		if divisor == 1 {
			switch remain {
			case 1, 2, 3:
				for remain > 0 {
					result[index], index, remain = 'I', index+1, remain-1
				}
			case 4:
				result[index], index = 'I', index+1
				result[index], index = 'V', index+1
			case 5, 6, 7, 8:
				result[index], index = 'V', index+1
				for remain > 5 {
					result[index], index, remain = 'I', index+1, remain-1
				}
			case 9:
				result[index], index = 'I', index+1
				result[index], index = 'X', index+1
			}
		} else if divisor == 10 {
			switch remain {
			case 1, 2, 3:
				for remain > 0 {
					result[index], index, remain = 'X', index+1, remain-1
				}
			case 4:
				result[index], index = 'X', index+1
				result[index], index = 'L', index+1
			case 5, 6, 7, 8:
				result[index], index = 'L', index+1
				for remain > 5 {
					result[index], index, remain = 'X', index+1, remain-1
				}
			case 9:
				result[index], index = 'X', index+1
				result[index], index = 'C', index+1
			}
		} else if divisor == 100 {
			switch remain {
			case 1, 2, 3:
				for remain > 0 {
					result[index], index, remain = 'C', index+1, remain-1
				}
			case 4:
				result[index], index = 'C', index+1
				result[index], index = 'D', index+1
			case 5, 6, 7, 8:
				result[index], index = 'D', index+1
				for remain > 5 {
					result[index], index, remain = 'C', index+1, remain-1
				}
			case 9:
				result[index], index = 'C', index+1
				result[index], index = 'M', index+1
			}
		} else if divisor == 1000 {
			for remain > 0 {
				result[index], index, remain = 'M', index+1, remain-1
			}
		}
		divisor = divisor / 10
		if divisor > 0 {
			remain = num / divisor
		}
	}
	return string(result[0:index])
}

func intToRoman2(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	var buffer bytes.Buffer
	buffer.WriteString(M[num/1000])
	num = num % 1000
	buffer.WriteString(C[num/100])
	num = num % 100
	buffer.WriteString(X[num/10])
	num = num % 10
	buffer.WriteString(I[num])
	return buffer.String()
}

func intToRoman3(num int) string {
	s := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	var buffer bytes.Buffer
	divisor := 1000
	for i := 0; i < 4; i++ {
		buffer.WriteString(s[i][num/divisor])
		num, divisor = num%divisor, divisor/10
	}
	return buffer.String()
}
