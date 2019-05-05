package numislands

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 *
 * https://leetcode.com/problems/number-of-islands/description/
 *
 * algorithms
 * Medium (40.86%)
 * Total Accepted:    341.5K
 * Total Submissions: 828.4K
 * Testcase Example:  '[["1","1","1","1","0"],["1","1","0","1","0"],["1","1","0","0","0"],["0","0","0","0","0"]]'
 *
 * Given a 2d grid map of '1's (land) and '0's (water), count the number of
 * islands. An island is surrounded by water and is formed by connecting
 * adjacent lands horizontally or vertically. You may assume all four edges of
 * the grid are all surrounded by water.
 * 
 * Example 1:
 * 
 * 
 * Input:
 * 11110
 * 11010
 * 11000
 * 00000
 * 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input:
 * 11000
 * 11000
 * 00100
 * 00011
 * 
 * Output: 3
 * 
 */
var rows, columns int

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns = len(grid), len(grid[0])
	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == '1' {
				DFSMarking(grid, i, j)
				count++
			}
		}
	}
	return count
}

func DFSMarking(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i >= rows || j >= columns || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '0'
	DFSMarking(grid, i+1, j)
	DFSMarking(grid, i-1, j)
	DFSMarking(grid, i, j+1)
	DFSMarking(grid, i, j-1)
}
