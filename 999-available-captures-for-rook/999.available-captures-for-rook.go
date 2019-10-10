package available_captures_for_rook

/*
 * @lc app=leetcode id=999 lang=golang
 *
 * [999] Available Captures for Rook
 *
 * https://leetcode.com/problems/available-captures-for-rook/description/
 *
 * algorithms
 * Easy (67.54%)
 * Total Accepted:    19.1K
 * Total Submissions: 29K
 * Testcase Example:  '[[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]'
 *
 * On an 8 x 8 chessboard, there is one white rook.  There also may be empty
 * squares, white bishops, and black pawns.  These are given as characters 'R',
 * '.', 'B', and 'p' respectively. Uppercase characters represent white pieces,
 * and lowercase characters represent black pieces.
 *
 * The rook moves as in the rules of Chess: it chooses one of four cardinal
 * directions (north, east, west, and south), then moves in that direction
 * until it chooses to stop, reaches the edge of the board, or captures an
 * opposite colored pawn by moving to the same square it occupies.  Also, rooks
 * cannot move into the same square as other friendly bishops.
 *
 * Return the number of pawns the rook can capture in one move.
 *
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input:
 * [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","R",".",".",".","p"],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
 * Output: 3
 * Explanation:
 * In this example the rook is able to capture all the pawns.
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input:
 * [[".",".",".",".",".",".",".","."],[".","p","p","p","p","p",".","."],[".","p","p","B","p","p",".","."],[".","p","B","R","B","p",".","."],[".","p","p","B","p","p",".","."],[".","p","p","p","p","p",".","."],[".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".","."]]
 * Output: 0
 * Explanation:
 * Bishops are blocking the rook to capture any pawn.
 *
 *
 * Example 3:
 *
 *
 *
 *
 * Input:
 * [[".",".",".",".",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".","p",".",".",".","."],["p","p",".","R",".","p","B","."],[".",".",".",".",".",".",".","."],[".",".",".","B",".",".",".","."],[".",".",".","p",".",".",".","."],[".",".",".",".",".",".",".","."]]
 * Output: 3
 * Explanation:
 * The rook can capture the pawns at positions b5, d6 and f5.
 *
 *
 *
 *
 * Note:
 *
 *
 * board.length == board[i].length == 8
 * board[i][j] is either 'R', '.', 'B', or 'p'
 * There is exactly one cell with board[i][j] == 'R'
 *
 *
 */
func numRookCaptures(board [][]byte) int {
	row, column := findR(board)
	ans := 0
	var ch byte
	for up := row - 1; up >= 0; up-- {
		ch = board[up][column]
		if ch == 'p' {
			ans++
			break
		} else if ch == 'R' || ch == 'B' {
			break
		}
	}
	for down := row + 1; down < 8; down++ {
		ch = board[down][column]
		if ch == 'p' {
			ans++
			break
		} else if ch == 'R' || ch == 'B' {
			break
		}
	}
	for right := column + 1; right < 8; right++ {
		ch = board[row][right]
		if ch == 'p' {
			ans++
			break
		} else if ch == 'R' || ch == 'B' {
			break
		}
	}
	for left := column - 1; left >= 0; left-- {
		ch = board[row][left]
		if ch == 'p' {
			ans++
			break
		} else if ch == 'R' || ch == 'B' {
			break
		}
	}
	return ans
}

func findR(board [][]byte) (int, int) {
	for column := 0; column < 8; column++ {
		for row := 0; row < 8; row++ {
			if board[row][column] == 'R' {
				return row, column
			}
		}
	}
	return -1, -1
}
