package fractiontodecimal

import (
	"bytes"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=166 lang=golang
 *
 * [166] Fraction to Recurring Decimal
 *
 * https://leetcode.com/problems/fraction-to-recurring-decimal/description/
 *
 * algorithms
 * Medium (19.31%)
 * Total Accepted:    86.2K
 * Total Submissions: 445K
 * Testcase Example:  '1\n2'
 *
 * Given two integers representing the numerator and denominator of a fraction,
 * return the fraction in string format.
 * 
 * If the fractional part is repeating, enclose the repeating part in
 * parentheses.
 * 
 * Example 1:
 * 
 * 
 * Input: numerator = 1, denominator = 2
 * Output: "0.5"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: numerator = 2, denominator = 1
 * Output: "2"
 * 
 * Example 3:
 * 
 * 
 * Input: numerator = 2, denominator = 3
 * Output: "0.(6)"
 * 
 * 
 */
var dotPositionMap map[int64]int
var numeratorMap map[int64]string
var dotPosition int

func fractionToDecimal(numerator int, denominator int) string {
	dotPositionMap = make(map[int64]int)
	numeratorMap = make(map[int64]string)

	var numeratorL = int64(numerator)
	var denominatorL = int64(denominator)
	var negativeFlag bool
	if numeratorL < 0 && denominatorL > 0 {
		numeratorL = -numeratorL
		negativeFlag = true
	} else if numeratorL < 0 && denominatorL < 0 {
		numeratorL = -numeratorL
		denominatorL = -denominatorL
	} else if numeratorL > 0 && denominatorL < 0 {
		denominatorL = - denominatorL
		negativeFlag = true
	}

	intPart := strconv.Itoa((int)(numeratorL/denominatorL)) + "."
	if negativeFlag {
		intPart = "-" + intPart
	}
	remain := numeratorL % denominatorL
	result := intPart + fractionToDecimalWithoutDot(remain*10, denominatorL, len(intPart))
	if strings.HasSuffix(result, ".0") {
		result = substring(result, 0, len(result)-2)
	} else if strings.HasSuffix(result, ")") {
		result = result[:len(result)-3]
		result = substring(result, 0, dotPosition) + "(" + substring(result, dotPosition, len(result)) + ")"
	}
	return result
}

func substring(s string, start, end int) string {
	var buffer bytes.Buffer
	for i := start; i < end; i++ {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}

func fractionToDecimalWithoutDot(numerator int64, denominator int64, position int) string {
	if numera, ok := numeratorMap[numerator]; ok {
		dotPosition = dotPositionMap[numerator]
		return "(" + numera + ")"
	}

	intPart := numerator / denominator
	numeratorMap[numerator] = strconv.Itoa(int(intPart))
	dotPositionMap[numerator] = position

	remains := numerator % denominator
	if remains == 0 {
		return numeratorMap[numerator]
	}
	return strconv.Itoa(int(intPart)) + fractionToDecimalWithoutDot(remains*10, denominator, position+1)

}
