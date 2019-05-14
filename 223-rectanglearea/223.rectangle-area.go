package computearea

/*
 * @lc app=leetcode id=223 lang=golang
 *
 * [223] Rectangle Area
 *
 * https://leetcode.com/problems/rectangle-area/description/
 *
 * algorithms
 * Medium (35.65%)
 * Total Accepted:    86.6K
 * Total Submissions: 242K
 * Testcase Example:  '-3\n0\n3\n4\n0\n-1\n9\n2'
 *
 * Find the total area covered by two rectilinear rectangles in a 2D plane.
 *
 * Each rectangle is defined by its bottom left corner and top right corner as
 * shown in the figure.
 *
 *
 *
 * Example:
 *
 *
 * Input: A = -3, B = 0, C = 3, D = 4, E = 0, F = -1, G = 9, H = 2
 * Output: 45
 *
 * Note:
 *
 * Assume that the total area is never beyond the maximum possible value of
 * int.
 *
 */
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	cross := crossArea(A, B, C, D, E, F, G, H)
	return (D-B)*(C-A) + (H-F)*(G-E) - cross
}

func crossArea(x1, y1, x2, y2, x3, y3, x4, y4 int) int {
	if y4 <= y1 || y3 >= y2 || x3 >= x2 || x1 >= x4 {
		return 0
	} else {
		xDis := distance(y1, y2, y3, y4)
		yDis := distance(x1, x2, x3, x4)
		return xDis * yDis
	}
}

func distance(y1, y2, y3, y4 int) int {
	if y2 >= y4 && y3 >= y1 {
		return y4 - y3
	} else if y2 <= y4 && y3 <= y1 {
		return y2 - y1
	} else {
		dy1 := abs(y2 - y3)
		dy2 := abs(y4 - y1)
		if dy1 > dy2 {
			return dy2
		} else {
			return dy1
		}
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
