package uniquepathswithobstacles

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 *
 * https://leetcode.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (33.30%)
 * Total Accepted:    191.7K
 * Total Submissions: 575.4K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * A robot is located at the top-left corner of a m x n grid (marked 'Start' in
 * the diagram below).
 *
 * The robot can only move either down or right at any point in time. The robot
 * is trying to reach the bottom-right corner of the grid (marked 'Finish' in
 * the diagram below).
 *
 * Now consider if some obstacles are added to the grids. How many unique paths
 * would there be?
 *
 *
 *
 * An obstacle and empty space is marked as 1 and 0 respectively in the grid.
 *
 * Note: m and n will be at most 100.ZL
 *
 * Example 1:
 *
 *
 * Input:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * Output: 2
 * Explanation:
 * There is one obstacle in the middle of the 3x3 grid above.
 * There are two ways to reach the bottom-right corner:
 * 1. Right -> Right -> Down -> Down
 * 2. Down -> Down -> Right -> Right
 *
 *
 */

var paths map[string]int

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	paths = make(map[string]int)
	width := len(obstacleGrid)
	height := len(obstacleGrid[0])
	if obstacleGrid[width-1][height-1] == 1 {
		return 0
	}
	return uniquePathsWithObstaclesIter(obstacleGrid, 0, 0, []int{width, height})
}

func uniquePathsWithObstaclesIter(board [][]int, startX, startY int, wh []int) int {
	key := strings.Join([]string{strconv.Itoa(startX), strconv.Itoa(startY)}, ",")
	if v, ok := paths[key]; ok {
		return v
	} else {
		if startX == wh[0]-1 && startY == wh[1]-1 {
			paths[key] = 1
			return 1
		} else if board[startX][startY] == 1 {
			paths[key] = 0
			return 0
		} else if startX == wh[0]-1 {
			ps := uniquePathsWithObstaclesIter(board, startX, startY+1, wh)
			paths[key] = ps
			return ps
		} else if startY == wh[1]-1 {
			ps := uniquePathsWithObstaclesIter(board, startX+1, startY, wh)
			paths[key] = ps
			return ps
		} else {
			ps := uniquePathsWithObstaclesIter(board, startX+1, startY, wh) + uniquePathsWithObstaclesIter(board, startX, startY+1, wh)
			paths[key] = ps
			return ps
		}
	}
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	columns := len(obstacleGrid[0])
	if obstacleGrid[rows-1][columns-1] == 1 {
		return 0
	}
	board := make([][]int, rows)
	for i := 0; i < rows; i++ {
		board[i] = make([]int, columns)
	}
	for i := rows - 1; i >= 0; i-- {
		for j := columns - 1; j >= 0; j-- {
			if obstacleGrid[i][j] == 1 {
				board[i][j] = 0
			} else {
				if i+1 < rows && j+1 < columns {
					board[i][j] = board[i][j+1] + board[i+1][j]
				} else if i+1 < rows {
					board[i][j] = board[i+1][j]
				} else if j+1 < columns {
					board[i][j] = board[i][j+1]
				} else {
					board[i][j] = 1
				}
			}
		}
	}
	return board[0][0]
}
