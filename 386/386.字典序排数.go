package main

import (
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=386 lang=golang
 *
 * [386] 字典序排数
 */

type NS []int

var keyAndSeparateMap map[int]string

func (ns NS) Len() int {
	return len(ns)
}

func (ns NS) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns NS) Less(i, j int) bool {
	vI, vJ := ns[i], ns[j]
	var iString, jString string
	var exist bool
	if iString, exist = keyAndSeparateMap[vI]; !exist {
		iString = strconv.Itoa(vI)
		keyAndSeparateMap[vI] = iString
	}
	if jString, exist = keyAndSeparateMap[vJ]; !exist {
		jString = strconv.Itoa(vJ)
		keyAndSeparateMap[vJ] = jString
	}
	for i := 0; i < len(iString) && i < len(jString); i++ {
		if iString[i] > jString[i] {
			return false
		} else if iString[i] < jString[i] {
			return true
		}
	}
	return len(iString) < len(jString)
}

// @lc code=start
func lexicalOrder(n int) []int {
	if keyAndSeparateMap == nil {
		keyAndSeparateMap = make(map[int]string)
	}
	numbers := make(NS, n, n)
	for i := 1; i <= n; i++ {
		numbers[i-1] = i
	}
	sort.Sort(numbers)
	return numbers
}

// @lc code=end
