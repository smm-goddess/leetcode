package smallest

import (
	"math"
)

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 *
 * https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/description/
 *
 * algorithms
 * Medium (48.89%)
 * Total Accepted:    109.6K
 * Total Submissions: 221.5K
 * Testcase Example:  '[[1,5,9],[10,11,13],[12,13,15]]\n8'
 *
 * Given a n x n matrix where each of the rows and columns are sorted in
 * ascending order, find the kth smallest element in the matrix.
 *
 *
 * Note that it is the kth smallest element in the sorted order, not the kth
 * distinct element.
 *
 *
 * Example:
 *
 * matrix = [
 * ⁠  [ 1,  5,  9],
 * ⁠  [10, 11, 13],
 * ⁠  [12, 13, 15]
 * ],
 * k = 8,
 *
 * return 13.
 *
 *
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ n2.
 */
//func kthSmallest(matrix [][]int, k int) int {
//	size := len(matrix)
//	indexes := make([]int, size, size)
//	smallestIndex, smallestValue := 0, 0
//	for k > 0 {
//		smallestValue = math.MaxInt32
//		for i := 0; i < size; i++ {
//			if indexes[i] < size {
//				if matrix[i][indexes[i]] < smallestValue {
//					smallestValue = matrix[i][indexes[i]]
//					smallestIndex = i
//				}
//			}
//		}
//		indexes[smallestIndex]++
//		k--
//	}
//	return smallestValue
//}

func kthSmallest(matrix [][]int, k int) int {
	size := len(matrix)
	indexes := make([]int, size, size)
	heap, value := make([]int, size, size), make([]int, size, size)
	for i := range heap {
		heap[i] = i
		value[i] = matrix[i][0]
	}
	heapify(heap, value, size)
	for k > 1 {
		indexes[heap[0]] = indexes[heap[0]] + 1
		if indexes[heap[0]] >= size {
			value[0] = math.MaxInt32
		} else {
			value[0] = matrix[heap[0]][indexes[heap[0]]]
		}
		heapify(heap, value, size)
		k--
	}
	return matrix[heap[0]][indexes[heap[0]]]
}

func heapify(heap, values []int, size int) {
	var leftSon, rightSon int
	for i := (size >> 1) - 1; i >= 0; i-- {
		leftSon = 2*i + 1
		rightSon = 2*i + 2
		if leftSon < size && values[leftSon] < values[i] {
			values[leftSon], values[i] = values[i], values[leftSon]
			heap[leftSon], heap[i] = heap[i], heap[leftSon]
		}
		if rightSon < size && values[rightSon] < values[i] {
			values[rightSon], values[i] = values[i], values[rightSon]
			heap[rightSon], heap[i] = heap[i], heap[rightSon]
		}
	}
}
