package searchmatrix

/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 *
 * https://leetcode.com/problems/search-a-2d-matrix/description/
 *
 * algorithms
 * Medium (34.75%)
 * Total Accepted:    215.1K
 * Total Submissions: 618.9K
 * Testcase Example:  '[[1,3,5,7],[10,11,16,20],[23,30,34,50]]\n3'
 *
 * Write an efficient algorithm that searches for a value in an m x n matrix.
 * This matrix has the following properties:
 *
 *
 * Integers in each row are sorted from left to right.
 * The first integer of each row is greater than the last integer of the
 * previous row.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 3
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input:
 * matrix = [
 * ⁠ [1,   3,  5,  7],
 * ⁠ [10, 11, 16, 20],
 * ⁠ [23, 30, 34, 50]
 * ]
 * target = 13
 * Output: false
 *
 */
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])
	if column == 0 || target < matrix[0][0] || target > matrix[row-1][column-1] {
		return false
	}
	first := make([]int, row)
	for i := 0; i < row; i++ {
		first[i] = matrix[i][0]
	}
	startRow := searchIndex(first, target)
	return binarySearch(matrix[startRow], target)
}

func binarySearch(row []int, target int) bool {
	start, end := 0, len(row)-1
	middle := start + (end-start)/2
	for end >= start {
		if row[start] == target || row[end] == target || row[middle] == target {
			return true
		} else if row[middle] > target {
			end = middle
			start = start + 1
			middle = start + (end-start)/2
		} else {
			start = middle
			end = end - 1
			middle = start + (end-start)/2
		}
	}
	return false
}

func searchIndex(column []int, target int) int {
	start, end := 0, len(column)-1
	middle := start + (end-start)/2
	if column[end] <= target {
		return end
	}
	for end-start > 1 {
		if column[start] == target {
			return start
		} else if column[end] == target {
			return end
		} else if column[middle] == target {
			return middle
		} else if column[middle] > target {
			end = middle
			middle = start + (end-start)/2
		} else if column[middle] < target {
			start = middle
			middle = start + (end-start)/2
		}
	}
	return start
}
