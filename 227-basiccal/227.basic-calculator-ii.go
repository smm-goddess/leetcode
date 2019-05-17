package calculator

import (
	"strings"
)

/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 *
 * https://leetcode.com/problems/basic-calculator-ii/description/
 *
 * algorithms
 * Medium (32.97%)
 * Total Accepted:    107.5K
 * Total Submissions: 322.6K
 * Testcase Example:  '"3+2*2"'
 *
 * Implement a basic calculator to evaluate a simple expression string.
 *
 * The expression string contains only non-negative integers, +, -, *, /
 * operators and empty spaces  . The integer division should truncate toward
 * zero.
 *
 * Example 1:
 *
 *
 * Input: "3+2*2"
 * Output: 7
 *
 *
 * Example 2:
 *
 *
 * Input: " 3/2 "
 * Output: 1
 *
 * Example 3:
 *
 *
 * Input: " 3+5 / 2 "
 * Output: 5
 *
 *
 * Note:
 *
 *
 * You may assume that the given expression is always valid.
 * Do not use the eval built-in library function.
 *
 *
 */
func calculate(s string) int {
	s = strings.Replace(s, " ", "", -1)
	operatorStack := make([]byte, 0)
	numberStash := make([]byte, 0)
	numberStack := make([]int, 0)
	for _, c := range []byte(s) {
		switch c {
		case '+', '-', '*', '/':
			numberStack = append(numberStack, transformToInt(numberStash))
			numberStash = numberStash[0:0]
			operatorStack = append(operatorStack, c)
		default:
			numberStash = append(numberStash, c)
		}
	}
	if len(numberStash) > 0 {
		numberStack = append(numberStack, transformToInt(numberStash))
	}
	filtered := make([]int, 0)
	for i, operator := range operatorStack {
		if operator == '*' || operator == '/' {
			filtered = append(filtered, i)
			if operator == '*' {
				numberStack[i+1] = numberStack[i] * numberStack[i+1]
			} else {
				numberStack[i+1] = numberStack[i] / numberStack[i+1]
			}
		}
	}
	for j := len(filtered) - 1; j >= 0; j-- {
		operatorStack = append(append([]byte{}, operatorStack[:filtered[j]]...), operatorStack[filtered[j]+1:]...)
		numberStack = append(append([]int{}, numberStack[:filtered[j]]...), numberStack[filtered[j]+1:]...)
	}
	for i, operator := range operatorStack {
		if operator == '+' {
			numberStack[i+1] = numberStack[i] + numberStack[i+1]
		} else {
			numberStack[i+1] = numberStack[i] - numberStack[i+1]
		}
	}
	return numberStack[len(numberStack)-1]
}

func transformToInt(bs []byte) int {
	if len(bs) == 0 {
		return -1
	}
	answer, bits := 0, 1
	for i := len(bs) - 1; i >= 0; i-- {
		answer += int(bs[i]-'0') * bits
		bits *= 10
	}
	return answer
}
