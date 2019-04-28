package solve

/*
 * @lc app=leetcode id=130 lang=golang
 *
 * [130] Surrounded Regions
 *
 * https://leetcode.com/problems/surrounded-regions/description/
 *
 * algorithms
 * Medium (22.41%)
 * Total Accepted:    140.5K
 * Total Submissions: 622.8K
 * Testcase Example:  '[["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]'
 *
 * Given a 2D board containing 'X' and 'O' (the letter O), capture all regions
 * surrounded by 'X'.
 *
 * A region is captured by flipping all 'O's into 'X's in that surrounded
 * region.
 *
 * Example:
 *
 *
 * X X X X
 * X O O X
 * X X O X
 * X O X X
 *
 *
 * After running your function, the board should be:
 *
 *
 * X X X X
 * X X X X
 * X X X X
 * X O X X
 *
 *
 * Explanation:
 *
 * Surrounded regions shouldn’t be on the border, which means that any 'O' on
 * the border of the board are not flipped to 'X'. Any 'O' that is not on the
 * border and it is not connected to an 'O' on the border will be flipped to
 * 'X'. Two cells are connected if they are adjacent cells connected
 * horizontally or vertically.
 *
 */
func solve(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	rows := len(board)
	columns := len(board[0])
	if rows == 1 || columns == 1 {
		return
	}
	todo := make([]point, 0)
	solved := make([]point, 0)
	for col := 0; col < columns; col++ {
		if board[0][col] == 'O' {
			todo = append(todo, point{0, col})
		}
		if board[rows-1][col] == 'O' {
			todo = append(todo, point{rows - 1, col})
		}
	}
	for row := 1; row < rows-1; row++ {
		if board[row][0] == 'O' {
			todo = append(todo, point{row, 0})
		}
		if board[row][columns-1] == 'O' {
			todo = append(todo, point{row, columns - 1})
		}
	}
	var found bool
	for len(todo) > 0 {
		newTodo := make([]point, 0)
		for _, wannaDo := range todo {
			x, y := wannaDo.x, wannaDo.y
			if x-1 >= 0 && board[x-1][y] == 'O' {
				found = false
				for _, s := range solved {
					if s.x == x-1 && s.y == y {
						found = true
					}
				}
				if !found {
					newTodo = append(newTodo, point{x - 1, y})
				}
			}
			if y-1 >= 0 && board[x][y-1] == 'O' {
				found = false
				for _, s := range solved {
					if s.x == x && s.y == y-1 {
						found = true
					}
				}
				if !found {
					newTodo = append(newTodo, point{x, y - 1})
				}
			}
			if x+1 < rows && board[x+1][y] == 'O' {
				found = false
				for _, s := range solved {
					if s.x == x+1 && s.y == y {
						found = true
					}
				}
				if !found {
					newTodo = append(newTodo, point{x + 1, y})
				}
			}
			if y+1 < columns && board[x][y+1] == 'O' {
				found = false
				for _, s := range solved {
					if s.x == x && s.y == y+1 {
						found = true
					}
				}
				if !found {
					newTodo = append(newTodo, point{x, y + 1})
				}
			}
			solved = append(solved, point{wannaDo.x, wannaDo.y})
		}
		todo = newTodo
	}
	for i := 0; i < rows; i ++ {
		for j := 0; j < columns; j++ {
			board[i][j] = 'X'
		}
	}
	for i := range solved {
		board[solved[i].x][solved[i].y] = 'O'
	}
}

type point struct {
	x int
	y int
}
