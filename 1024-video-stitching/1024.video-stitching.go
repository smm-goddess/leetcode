package video_stitching

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=1024 lang=golang
 *
 * [1024] Video Stitching
 *
 * https://leetcode.com/problems/video-stitching/description/
 *
 * algorithms
 * Medium (47.54%)
 * Total Accepted:    13.8K
 * Total Submissions: 29.1K
 * Testcase Example:  '[[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]]\n10'
 *
 * You are given a series of video clips from a sporting event that lasted T
 * seconds.  These video clips can be overlapping with each other and have
 * varied lengths.
 *
 * Each video clip clips[i] is an interval: it starts at time clips[i][0] and
 * ends at time clips[i][1].  We can cut these clips into segments freely: for
 * example, a clip [0, 7] can be cut into segments [0, 1] + [1, 3] + [3, 7].
 *
 * Return the minimum number of clips needed so that we can cut the clips into
 * segments that cover the entire sporting event ([0, T]).  If the task is
 * impossible, return -1.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: clips = [[0,2],[4,6],[8,10],[1,9],[1,5],[5,9]], T = 10
 * Output: 3
 * Explanation:
 * We take the clips [0,2], [8,10], [1,9]; a total of 3 clips.
 * Then, we can reconstruct the sporting event as follows:
 * We cut [1,9] into segments [1,2] + [2,8] + [8,9].
 * Now we have segments [0,2] + [2,8] + [8,10] which cover the sporting event
 * [0, 10].
 *
 *
 * Example 2:
 *
 *
 * Input: clips = [[0,1],[1,2]], T = 5
 * Output: -1
 * Explanation:
 * We can't cover [0,5] with only [0,1] and [0,2].
 *
 *
 * Example 3:
 *
 *
 * Input: clips =
 * [[0,1],[6,8],[0,2],[5,6],[0,4],[0,3],[6,7],[1,3],[4,7],[1,4],[2,5],[2,6],[3,4],[4,5],[5,7],[6,9]],
 * T = 9
 * Output: 3
 * Explanation:
 * We can take clips [0,4], [4,7], and [6,9].
 *
 *
 * Example 4:
 *
 *
 * Input: clips = [[0,4],[2,8]], T = 5
 * Output: 2
 * Explanation:
 * Notice you can have extra video after the event ends.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= clips.length <= 100
 * 0 <= clips[i][0], clips[i][1] <= 100
 * 0 <= T <= 100
 *
 *
 */

type IntRange [][]int

func (this IntRange) Len() int {
	return len(this)
}

func (this IntRange) Less(i, j int) bool {
	if this[i][0] == this[j][0] {
		return this[i][1] > this[j][1]
	} else {
		return this[i][0] < this[j][0]
	}
}

func (this IntRange) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

var tVal int

func videoStitching(clips [][]int, T int) int {
	sort.Sort(IntRange(clips))
	tVal = T
	min := math.MaxInt32
	stitching(clips, 0, 0, &min)
	if min == math.MaxInt32 {
		return -1
	} else {
		return min
	}
}

func stitching(clips [][]int, endPoint, count int, min *int) {
	if endPoint >= tVal && endPoint != 0 && count < *min {
		*min = count
	} else if count >= *min {
		return
	} else {
		for index, clip := range clips {
			if clip[0] <= endPoint && clip[1] > endPoint {
				stitching(clips[index:], clip[1], count+1, min)
			}
		}
	}
}
