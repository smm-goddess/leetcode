//package ladderlength

//import "fmt"

/*
 * @lc app=leetcode id=127 lang=golang
 *
 * [127] Word Ladder
 *
 * https://leetcode.com/problems/word-ladder/description/
 *
 * algorithms
 * Medium (23.44%)
 * Total Accepted:    248.9K
 * Total Submissions: 1.1M
 * Testcase Example:  '"hit"\n"cog"\n["hot","dot","dog","lot","log","cog"]'
 *
 * Given two words (beginWord and endWord), and a dictionary's word list, find
 * the length of shortest transformation sequence from beginWord to endWord,
 * such that:
 *
 *
 * Only one letter can be changed at a time.
 * Each transformed word must exist in the word list. Note that beginWord is
 * not a transformed word.
 *
 *
 * Note:
 *
 *
 * Return 0 if there is no such transformation sequence.
 * All words have the same length.
 * All words contain only lowercase alphabetic characters.
 * You may assume no duplicates in the word list.
 * You may assume beginWord and endWord are non-empty and are not the same.
 *
 *
 * Example 1:
 *
 *
 * Input:
 * beginWord = "hit",
 * endWord = "cog",
 * wordList = ["hot","dot","dog","lot","log","cog"]
 *
 * Output: 5
 *
 * Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" ->
 * "dog" -> "cog",
 * return its length 5.
 *
 *
 * Example 2:
 *
 *
 * Input:
 * beginWord = "hit"
 * endWord = "cog"
 * wordList = ["hot","dot","dog","lot","log"]
 *
 * Output: 0
 *
 * Explanation: The endWord "cog" is not in wordList, therefore no possible
 * transformation.
 *
 *
 *
 *
 *
 */
func ladderLength(beginWord string, endWord string, wordList []string) int {
	distanceQueue := make(map[string]int)
	distanceQueueCopy := make(map[string]int)
	distanceQueue[beginWord] = 1
	for len(distanceQueue) > 0 {
		for k, v := range distanceQueue {
			if k == endWord {
				return v
			}
			neighbors := make([]int, 0)
			for i, s := range wordList {
				if neighbor(k, s) {
					neighbors = append(neighbors, i)
				}
			}
			for i := len(neighbors) - 1; i >= 0; i-- {
				distanceQueueCopy[wordList[neighbors[i]]] = v + 1
				wordList = append(append([]string{}, wordList[:neighbors[i]]...), wordList[neighbors[i]+1:]...)
			}
		}
		distanceQueue, distanceQueueCopy = distanceQueueCopy, distanceQueue
		distanceQueueCopy = make(map[string]int)
	}
	return 0
}

func distance(a, b string) (dist int) {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return
}

func neighbor(a, b string) bool {
	return distance(a, b) == 1
}
