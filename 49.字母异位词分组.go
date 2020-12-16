/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	helper := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, s := range str {
			cnt[s-'a']++
		}
		helper[cnt] = append(helper[cnt], str)
	}

	result := [][]string{}

	for _, v := range helper {
		result = append(result, v)
	}
	return result
}

// @lc code=end

