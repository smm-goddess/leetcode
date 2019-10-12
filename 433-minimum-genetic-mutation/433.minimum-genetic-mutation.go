package minimum_genetic_mutation

import (
	"math"
)

/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 *
 * https://leetcode.com/problems/minimum-genetic-mutation/description/
 *
 * algorithms
 * Medium (37.67%)
 * Total Accepted:    24.4K
 * Total Submissions: 62K
 * Testcase Example:  '"AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]'
 *
 * A gene string can be represented by an 8-character long string, with choices
 * from "A", "C", "G", "T".
 *
 * Suppose we need to investigate about a mutation (mutation from "start" to
 * "end"), where ONE mutation is defined as ONE single character changed in the
 * gene string.
 *
 * For example, "AACCGGTT" -> "AACCGGTA" is 1 mutation.
 *
 * Also, there is a given gene "bank", which records all the valid gene
 * mutations. A gene must be in the bank to make it a valid gene string.
 *
 * Now, given 3 things - start, end, bank, your task is to determine what is
 * the minimum number of mutations needed to mutate from "start" to "end". If
 * there is no such a mutation, return -1.
 *
 * Note:
 *
 *
 * Starting point is assumed to be valid, so it might not be included in the
 * bank.
 * If multiple mutations are needed, all mutations during in the sequence must
 * be valid.
 * You may assume start and end string is not the same.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * start: "AACCGGTT"
 * end:   "AACCGGTA"
 * bank: ["AACCGGTA"]
 * return: 1
 *
 *
 *
 *
 *
 * Example 2:
 *
 *
 * start: "AACCGGTT"
 * end:   "AAACGGTA"
 * bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]
 *
 * return: 2
 *
 *
 *
 *
 * Example 3:
 *
 *
 * start: "AAAAACCC"
 * end:   "AACCCCCC"
 * bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]
 *
 * return: 3
 *
 *
 *
 *
 */
func minMutation(start string, end string, bank []string) int {
	startEndMap, distance, weightMap := make(map[string]int), make(map[string]int), make(map[string]map[string]int)
	for index := 0; index < len(bank); index++ {
		distance[bank[index]] = geneDistance(start, bank[index])
		for i := 0; i < len(bank); i++ {
			if m, ok := weightMap[bank[index]]; ok {
				m[bank[i]] = geneDistance(bank[index], bank[i])
			} else {
				weightMap[bank[index]] = make(map[string]int)
				weightMap[bank[index]][bank[i]] = geneDistance(bank[index], bank[i])
			}
		}
	}
	minKey, minDistance := "", math.MaxInt32
	for len(distance) > 0 {
		for key, value := range distance {
			if value <= minDistance {
				minKey, minDistance = key, value
			}
		}
		startEndMap[minKey] = minDistance
		for key, dist := range distance {
			if weightMap[key][minKey] != math.MaxInt32 || dist != math.MaxInt32 {
				distance[key] = minInt(dist, minDistance+weightMap[minKey][key])
			} else {
				distance[key] = math.MaxInt32
			}
		}
		delete(distance, minKey)
		minKey, minDistance = "", math.MaxInt32
	}
	if dist, ok := startEndMap[end]; ok && dist != math.MaxInt32 {
		return dist
	} else {
		return -1
	}
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func geneDistance(start, end string) int {
	diff := 0
	for index := 0; index < len(start); index++ {
		if start[index] != end[index] {
			diff++
			if diff > 1 {
				return math.MaxInt32
			}
		}
	}
	return diff
}
