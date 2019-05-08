package findorder

/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 *
 * https://leetcode.com/problems/course-schedule-ii/description/
 *
 * algorithms
 * Medium (34.19%)
 * Total Accepted:    143.4K
 * Total Submissions: 415.2K
 * Testcase Example:  '2\n[[1,0]]'
 *
 * There are a total of n courses you have to take, labeled from 0 to n-1.
 *
 * Some courses may have prerequisites, for example to take course 0 you have
 * to first take course 1, which is expressed as a pair: [0,1]
 *
 * Given the total number of courses and a list of prerequisite pairs, return
 * the ordering of courses you should take to finish all courses.
 *
 * There may be multiple correct orders, you just need to return one of them.
 * If it is impossible to finish all courses, return an empty array.
 *
 * Example 1:
 *
 *
 * Input: 2, [[1,0]]
 * Output: [0,1]
 * Explanation: There are a total of 2 courses to take. To take course 1 you
 * should have finished
 * course 0. So the correct course order is [0,1] .
 *
 * Example 2:
 *
 *
 * Input: 4, [[1,0],[2,0],[3,1],[3,2]]
 * Output: [0,1,2,3] or [0,2,1,3]
 * Explanation: There are a total of 4 courses to take. To take course 3 you
 * should have finished both
 * ⁠            courses 1 and 2. Both courses 1 and 2 should be taken after you
 * finished course 0.
 * So one correct course order is [0,1,2,3]. Another correct ordering is
 * [0,2,1,3] .
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
func findOrder(numCourses int, prerequisites [][]int) []int {
	if prerequisites == nil || len(prerequisites) == 0 {
		result := make([]int, numCourses, numCourses)
		for i := 0; i < numCourses; i++ {
			result[i] = i
		}
		return result
	} else if len(prerequisites) == 1 {
		prerequisite := prerequisites[0]
		result := make([]int, numCourses, numCourses)
		for i := 0; i < numCourses; i++ {
			result[i] = i
		}
		if prerequisite[0] < prerequisite[1] {
			result[prerequisite[0]], result[prerequisite[1]] = result[prerequisite[1]], result[prerequisite[0]]
		}
		return result
	}
	var result []int
	inDegreeZero := make([]bool, numCourses, numCourses)
	for i := range inDegreeZero {
		inDegreeZero[i] = true
	}
	for _, v := range prerequisites {
		inDegreeZero[v[1]] = false
	}
	for i, v := range inDegreeZero {
		if v {
			result = append(result, i)
		}
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
					result = append(result, v)
					inDegreeZero[v] = true
					newInDegreeZeroNumbers = append(newInDegreeZeroNumbers, v)
				}
			}
		}
		newInDegreeZeroNumbers, inDegreeZeroNumbers = inDegreeZeroNumbers, newInDegreeZeroNumbers
	}
	for _, v := range inDegreeZero {
		if !v {
			return []int{}
		}
	}
	r := make([]int, len(result), len(result))
	for i := len(result) - 1; i >= 0; i-- {
		r[len(result)-1-i] = result[i]
	}
	return r
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
