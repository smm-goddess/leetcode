package canfinish

/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 *
 * https://leetcode.com/problems/course-schedule/description/
 *
 * algorithms
 * Medium (37.19%)
 * Total Accepted:    208.1K
 * Total Submissions: 554.5K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, is it
 * possible for you to finish all courses?
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: true
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0. So it is possible.
 *
 * Example 2:
 *
 *
 * Input: 2, [[1,0],[0,1]]
 * Output: false
 * Explanation: There are a total of 2 courses to take.
 * To take course 1 you should have finished course 0, and to take course 0 you
 * should
 * also have finished course 1. So it is impossible.
 *
 *
 * Note:
 *
 *
 * The input prerequisites is a graph represented by a list of edges, not
 * adjacency matrices. Read more about how a graph is represented.
 * You may assume that there are no duplicate edges in the input
 * prerequisites.
 *
 *
 */
// kahn算法
func canFinish(numCourses int, prerequisites [][]int) bool {
	if prerequisites == nil || len(prerequisites) < 2 {
		return true
	}
	inDegreeZero := make([]bool, numCourses, numCourses)
	for i := range inDegreeZero {
		inDegreeZero[i] = true
	}
	for _, v := range prerequisites {
		inDegreeZero[v[1]] = false
	}
	inDegreeZeroNumbers := make([]int, 0)
	newInDegreeZeroNumbers := make([]int, 0)
	for i, v := range inDegreeZero {
		if v {
			inDegreeZeroNumbers = append(inDegreeZeroNumbers, i)
		}
	}
	for len(inDegreeZeroNumbers) > 0 {
		newInDegreeZeroNumbers = newInDegreeZeroNumbers[0:0]
		for _, n := range inDegreeZeroNumbers {
			endPoints := deleteEdge(&prerequisites, n)
			for _, v := range endPoints {
				if isZeroInDegree(prerequisites, v) {
					inDegreeZero[v] = true
					newInDegreeZeroNumbers = append(newInDegreeZeroNumbers, v)
				}
			}
		}
		newInDegreeZeroNumbers, inDegreeZeroNumbers = inDegreeZeroNumbers, newInDegreeZeroNumbers
	}
	for _, v := range inDegreeZero {
		if !v {
			return false
		}
	}
	return true
}

func isZeroInDegree(prerequisites [][]int, number int) bool {
	for _, v := range prerequisites {
		if v[1] == number {
			return false
		}
	}
	return true
}

func deleteEdge(prerequisites *[][]int, start int) (endPoint []int) {
	mark := make([]int, 0)
	for i, v := range *prerequisites {
		if v[0] == start {
			mark = append(mark, i)
		}
	}
	for i := len(mark) - 1; i >= 0; i-- {
		endPoint = append(endPoint, (*prerequisites)[mark[i]][1])
		*prerequisites = append(append([][]int{}, (*prerequisites)[:mark[i]]...), (*prerequisites)[mark[i]+1:]...)
	}
	return
}
