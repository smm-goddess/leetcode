package largestnumber

import (
	"bytes"
	"sort"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 *
 * https://leetcode.com/problems/largest-number/description/
 *
 * algorithms
 * Medium (25.48%)
 * Total Accepted:    126.6K
 * Total Submissions: 494K
 * Testcase Example:  '[10,2]'
 *
 * Given a list of non negative integers, arrange them such that they form the
 * largest number.
 * 
 * Example 1:
 * 
 * 
 * Input: [10,2]
 * Output: "210"
 * 
 * Example 2:
 * 
 * 
 * Input: [3,30,34,5,9]
 * Output: "9534330"
 * 
 * 
 * Note: The result may be very large, so you need to return a string instead
 * of an integer.
 * 
 */
func largestNumber(nums []int) string {
	sort.Sort(intSlice(nums))
	var buffer bytes.Buffer
	for _, num := range nums {
		buffer.WriteString(strconv.Itoa(num))
	}
	str := buffer.String()
	if strings.HasPrefix(str, "0") {
		return "0"
	} else {
		return str
	}
}

type intSlice []int

func (p intSlice) Len() int {
	return len(p)
}
func (p intSlice) Less(i, j int) bool {
	vi, vj, position := p[i], p[j], 0
	startI, startJ, li, lj := getNumberPosition(vi, position), getNumberPosition(vj, position), getNumberLength(vi), getNumberLength(vj)
	var iLonger bool
	if li < lj {
		for lj-li > 0 {
			vi = vi*10 + startI
			li ++
		}
	} else {
		iLonger = true
		for li-lj > 0 {
			vj = vj*10 + startJ
			lj ++
		}
	}
	if vi != vj {
		return vi > vj
	} else {
		for position < li {
			position++
			pi := getNumberPosition(vi, position)
			if pi > startI {
				return !iLonger
			} else if pi < startI {
				return iLonger
			}
		}
	}
	return false
}

func (p intSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func getNumberLength(num int) (length int) {
	for num > 0 {
		length = length + 1
		num = num / 10
	}
	return
}

func getNumberPosition(num, pos int) int {
	if num == 0 {
		return 0
	} else {
		var ints []int
		for num > 0 {
			ints = append(ints, num%10)
			num = num / 10
		}
		if pos >= len(ints) {
			return ints[len(ints)-1]
		} else {
			return ints[len(ints)-1-pos]
		}
	}
}
