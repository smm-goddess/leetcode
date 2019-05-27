package minheighttree

/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 *
 * https://leetcode.com/problems/minimum-height-trees/description/
 *
 * algorithms
 * Medium (29.94%)
 * Total Accepted:    64.1K
 * Total Submissions: 213K
 * Testcase Example:  '4\n[[1,0],[1,2],[1,3]]'
 *
 * For an undirected graph with tree characteristics, we can choose any node as
 * the root. The result graph is then a rooted tree. Among all possible rooted
 * trees, those with minimum height are called minimum height trees (MHTs).
 * Given such a graph, write a function to find all the MHTs and return a list
 * of their root labels.
 *
 * Format
 * The graph contains n nodes which are labeled from 0 to n - 1. You will be
 * given the number n and a list of undirected edges (each edge is a pair of
 * labels).
 *
 * You can assume that no duplicate edges will appear in edges. Since all edges
 * are undirected, [0, 1] is the same as [1, 0] and thus will not appear
 * together in edges.
 *
 * Example 1 :
 *
 *
 * Input: n = 4, edges = [[1, 0], [1, 2], [1, 3]]
 *
 * ⁠       0
 * ⁠       |
 * ⁠       1
 * ⁠      / \
 * ⁠     2   3
 *
 * Output: [1]
 *
 *
 * Example 2 :
 *
 *
 * Input: n = 6, edges = [[0, 3], [1, 3], [2, 3], [4, 3], [5, 4]]
 *
 * ⁠    0  1  2
 * ⁠     \ | /
 * ⁠       3
 * ⁠       |
 * ⁠       4
 * ⁠       |
 * ⁠       5
 *
 * Output: [3, 4]
 *
 * Note:
 *
 *
 * According to the definition of tree on Wikipedia: “a tree is an undirected
 * graph in which any two vertices are connected by exactly one path. In other
 * words, any connected graph without simple cycles is a tree.”
 * The height of a rooted tree is the number of edges on the longest downward
 * path between the root and a leaf.
 *
 *
 */
func findMinHeightTrees(n int, edges [][]int) []int {
	degreeCount := make([]int, n, n)
	leafNode := make([]bool, n, n)
	var cutoffIndex []int
	var newEdges [][]int
	for len(edges) > 1 {
		for _, edge := range edges {
			degreeCount[edge[0]] ++
			degreeCount[edge[1]] ++
		}
		for i, degree := range degreeCount {
			if degree == 1 {
				leafNode[i] = true
			} else {
				degreeCount[i] = 0
			}
		}
		for i, edge := range edges {
			if leafNode[edge[0]] || leafNode[edge[1]] {
				cutoffIndex = append(cutoffIndex, i)
			}
		}
		startIndex := 0
		for _, v := range cutoffIndex {
			newEdges = append(newEdges, edges[startIndex:v]...)
			startIndex = v + 1
		}
		newEdges = append(newEdges, edges[startIndex:]...)
		edges, newEdges = newEdges, edges
		cutoffIndex = cutoffIndex[0:0]
		newEdges = newEdges[0:0]
	}
	if len(edges) == 1 {
		return edges[0]
	} else {
		for i, v := range leafNode {
			if !v {
				return []int{i}
			}
		}
	}
	return nil
}

/*time limit
func findMinHeightTrees(n int, edges [][]int) []int {
	minHeight := math.MaxInt32
	var minHeightNode []int
	for i := 0; i < n; i++ {
		height := heightTree(n, i, edges)
		if height == minHeight {
			minHeightNode = append(minHeightNode, i)
		} else if height < minHeight {
			minHeightNode, minHeight = []int{i}, height
		}
	}
	return minHeightNode
}

func heightTree(total, root int, edges [][]int) (height int) {
	neighbors, visited := []int{root}, make([]bool, total, total)
	visited[root] = true
	var newNeighbors []int
	for len(neighbors) > 0 {
		for _, neighbor := range neighbors {
			for _, edge := range edges {
				if edge[0] == neighbor && !visited[edge[1]] {
					newNeighbors = append(newNeighbors, edge[1])
					visited[edge[1]] = true
				} else if edge[1] == neighbor && !visited[edge[0]] {
					newNeighbors = append(newNeighbors, edge[0])
					visited[edge[0]] = true
				}
			}
		}
		height++
		neighbors, newNeighbors = newNeighbors, neighbors
		newNeighbors = newNeighbors[0:0]
	}
	return
}
*/
