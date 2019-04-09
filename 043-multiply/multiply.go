package multiply

/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 *
 * https://leetcode.com/problems/multiply-strings/description/
 *
 * algorithms
 * Medium (30.25%)
 * Total Accepted:    191.4K
 * Total Submissions: 632.8K
 * Testcase Example:  '"2"\n"3"'
 *
 * Given two non-negative integers num1 and num2 represented as strings, return
 * the product of num1 and num2, also represented as a string.
 *
 * Example 1:
 *
 *
 * Input: num1 = "2", num2 = "3"
 * Output: "6"
 *
 * Example 2:
 *
 *
 * Input: num1 = "123", num2 = "456"
 * Output: "56088"
 *
 *
 * Note:
 *
 *
 * The length of both num1 and num2 is < 110.
 * Both num1 and num2 contain only digits 0-9.
 * Both num1 and num2 do not contain any leading zero, except the number 0
 * itself.
 * You must not use any built-in BigInteger library or convert the inputs to
 * integer directly.
 *
 *
 */
func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 || num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	base := make([]byte, len(num1))
	for i := 0; i < len(num1); i++ {
		base[len(num1)-i-1] = num1[i] - '0'
	}
	multi := make([][]byte, len(num2))
	for i := 0; i < len(num2); i++ {
		zeros := len(num2) - i - 1
		multiRow := make([]byte, zeros+len(num1))
		num := num2[i] - '0'
		for j := 0; j < len(num1); j++ {
			multiRow[j+zeros] = num * base[j]
		}
		multi[i] = multiRow
	}
	maxLength := 0
	for _, mul := range multi {
		if maxLength < len(mul) {
			maxLength = len(mul)
		}
	}
	var result []byte
	sum := 0
	for i := 0; i < maxLength; i++ {
		for _, v := range multi {
			if len(v) > i {
				sum += int(v[i])
			}
		}
		result = append([]byte{byte(sum % 10)}, result...)
		sum = sum / 10
	}
	for sum > 0 {
		result = append([]byte{byte(sum % 10)}, result...)
		sum = sum / 10
	}
	for i := range result {
		result[i] += '0'
	}
	return string(result)
}
