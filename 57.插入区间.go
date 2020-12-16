import "sort"

/*
 * @lc app=leetcode.cn id=57 lang=golang
 *
 * [57] 插入区间
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {

	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	result := make([][]int, 0)
	for i := 0; i < n; i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		for j := i + 1; j < n; j++ {
			if end >= intervals[j][0] {
				if end < intervals[j][1] {
					end = intervals[j][1]
				}
			} else {
				i = j - 1
				break
			}
			if j == n-1 {
				tmp := []int{start, end}
				result = append(result, tmp)
				return result
			}
		}
		tmp := []int{start, end}
		result = append(result, tmp)
	}
	return result

}

// @lc code=end

