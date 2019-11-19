package solve_equation

import "fmt"

/*
 * @lc app=leetcode id=640 lang=golang
 *
 * [640] Solve the Equation
 *
 * https://leetcode.com/problems/solve-the-equation/description/
 *
 * algorithms
 * Medium (40.95%)
 * Total Accepted:    20.6K
 * Total Submissions: 50.2K
 * Testcase Example:  '"x+5-3+x=6+x-2"'
 *
 *
 * Solve a given equation and return the value of x in the form of string
 * "x=#value". The equation contains only '+', '-' operation, the variable x
 * and its coefficient.
 *
 *
 *
 * If there is no solution for the equation, return "No solution".
 *
 *
 * If there are infinite solutions for the equation, return "Infinite
 * solutions".
 *
 *
 * If there is exactly one solution for the equation, we ensure that the value
 * of x is an integer.
 *
 *
 * Example 1:
 *
 * Input: "x+5-3+x=6+x-2"
 * Output: "x=2"
 *
 *
 *
 * Example 2:
 *
 * Input: "x=x"
 * Output: "Infinite solutions"
 *
 *
 *
 * Example 3:
 *
 * Input: "2x=x"
 * Output: "x=0"
 *
 *
 *
 * Example 4:
 *
 * Input: "2x+3x-6x=x+2"
 * Output: "x=-1"
 *
 *
 *
 * Example 5:
 *
 * Input: "x=x+2"
 * Output: "No solution"
 *
 *
 */
func solveEquation(equation string) string {
	var number, leftNumber, leftCoefficient, rightNumber, rightCoefficient int
	positive, left := 1, true
	for pos, c := range []byte(equation) {
		switch c {
		case '=':
			leftNumber += positive * number
			left, number, positive = false, 0, 1
		case '+':
			if left {
				leftNumber += positive * number
			} else {
				rightNumber += positive * number
			}
			positive, number = 1, 0
		case '-':
			if left {
				leftNumber += positive * number
			} else {
				rightNumber += positive * number
			}
			positive, number = -1, 0
		case 'x':
			if number == 0 {
				if pos >= 1 && equation[pos-1] == '0' {
					number = 0
				} else {
					number = 1
				}
			}
			if left {
				leftCoefficient += positive * number
			} else {
				rightCoefficient += positive * number
			}
			number = 0
		default:
			number = number*10 + int(c-'0')
		}
	}
	rightNumber += positive * number
	coefficient := leftCoefficient - rightCoefficient
	number = rightNumber - leftNumber
	if coefficient == 0 && number == 0 {
		return "Infinite solutions"
	} else if coefficient == 0 && number != 0 {
		return "No solution"
	} else {
		return fmt.Sprintf("x=%d", number/coefficient)
	}
}
