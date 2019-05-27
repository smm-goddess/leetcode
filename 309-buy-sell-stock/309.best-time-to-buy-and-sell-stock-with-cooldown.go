package maxprofit

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
 *
 * algorithms
 * Medium (43.69%)
 * Total Accepted:    90.1K
 * Total Submissions: 205K
 * Testcase Example:  '[1,2,3,0,2]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * Design an algorithm to find the maximum profit. You may complete as many
 * transactions as you like (ie, buy one and sell one share of the stock
 * multiple times) with the following restrictions:
 *
 *
 * You may not engage in multiple transactions at the same time (ie, you must
 * sell the stock before you buy again).
 * After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1
 * day)
 *
 *
 * Example:
 *
 *
 * Input: [1,2,3,0,2]
 * Output: 3
 * Explanation: transactions = [buy, sell, cooldown, buy, sell]
 *
 */
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	length := len(prices)
	buy, sell, cooldown := make([]int, length, length), make([]int, length, length), make([]int, length, length)
	buy[0] = -prices[0]
	for i := 1; i < length; i++ {
		cooldown[i] = sell[i-1]
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
		buy[i] = max(buy[i-1], cooldown[i-1]-prices[i])
	}
	return max(sell[length-1], cooldown[length-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
