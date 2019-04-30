package evalrpn

import "strconv"

/*
 * @lc app=leetcode id=150 lang=golang
 *
 * [150] Evaluate Reverse Polish Notation
 *
 * https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
 *
 * algorithms
 * Medium (31.71%)
 * Total Accepted:    158.2K
 * Total Submissions: 495.5K
 * Testcase Example:  '["2","1","+","3","*"]'
 *
 * Evaluate the value of an arithmetic expression in Reverse Polish Notation.
 * 
 * Valid operators are +, -, *, /. Each operand may be an integer or another
 * expression.
 * 
 * Note:
 * 
 * 
 * Division between two integers should truncate toward zero.
 * The given RPN expression is always valid. That means the expression would
 * always evaluate to a result and there won't be any divide by zero
 * operation.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: ["2", "1", "+", "3", "*"]
 * Output: 9
 * Explanation: ((2 + 1) * 3) = 9
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: ["4", "13", "5", "/", "+"]
 * Output: 6
 * Explanation: (4 + (13 / 5)) = 6
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
 * Output: 22
 * Explanation: 
 * ⁠ ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
 * = ((10 * (6 / (12 * -11))) + 17) + 5
 * = ((10 * (6 / -132)) + 17) + 5
 * = ((10 * 0) + 17) + 5
 * = (0 + 17) + 5
 * = 17 + 5
 * = 22
 * 
 * 
 */
func evalRPN(tokens []string) int {
	if tokens == nil || len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0, len(tokens))
	for _, t := range tokens {
		switch t {
		case "+":
			length := len(stack)
			stack[length-2] = stack[length-2] + stack[length-1]
			stack = stack[:length-1]
		case "-":
			length := len(stack)
			stack[length-2] = stack[length-2] - stack[length-1]
			stack = stack[:length-1]
		case "*":
			length := len(stack)
			stack[length-2] = stack[length-2] * stack[length-1]
			stack = stack[:length-1]
		case "/":
			length := len(stack)
			stack[length-2] = stack[length-2] / stack[length-1]
			stack = stack[:length-1]
		default:
			number, _ := strconv.Atoi(t)
			stack = append(stack, number)
		}
	}
	return stack[0]
}
