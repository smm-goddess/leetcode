package coinchange

import (
	"math"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 *
 * https://leetcode.com/problems/coin-change/description/
 *
 * algorithms
 * Medium (29.70%)
 * Total Accepted:    189.8K
 * Total Submissions: 629.5K
 * Testcase Example:  '[1,2,5]\n11'
 *
 * You are given coins of different denominations and a total amount of money
 * amount. Write a function to compute the fewest number of coins that you need
 * to make up that amount. If that amount of money cannot be made up by any
 * combination of the coins, return -1.
 *
 * Example 1:
 *
 *
 * Input: coins = [1, 2, 5], amount = 11
 * Output: 3
 * Explanation: 11 = 5 + 5 + 1
 *
 * Example 2:
 *
 *
 * Input: coins = [2], amount = 3
 * Output: -1
 *
 *
 * Note:
 * You may assume that you have an infinite number of each kind of coin.
 *
 */
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	} else if coins == nil || len(coins) == 0 || amount < 0 {
		return -1
	}
	dp := make([]int, amount+1, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for _, c := range coins {
		if c <= amount {
			dp[c] = 1
		}
	}
	i, j, k := 1, 1, 0
	for ; i <= amount; i++ {
		if dp[i] != 1 {
			j = 1
			for ; j <= i>>1; j++ {
				k = i - j
				if dp[j] != math.MaxInt32 && dp[k] != math.MaxInt32 {
					dp[i] = min(dp[i], dp[j]+dp[k])
				}
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
