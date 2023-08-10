/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
func maxPoints(points [][]int) int {

	if len(points) <= 1 {
		return len(points)
	}

	var max = 0
	for i := 0; i < len(points); i++ {
		var maxMap = make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			a := float64(points[j][1]-points[i][1]) / float64(points[j][0]-points[i][0])
			b := float64(points[j][0]*points[i][1]-points[i][0]*points[j][1]) / float64(points[j][0]-points[i][0])
			if b == 0 {
				b = 0
			}
			if math.IsInf(b, 0) {
				b = float64(math.Inf(0))
			}
			if math.IsInf(a, 0) {
				a = float64(math.Inf(0))
			}
			if a == 0 {
				a = 0
			}
			// fmt.Printf("%v,%v,%f:%f\n", points[i], points[j], a, b)
			if _, ok := maxMap[fmt.Sprintf("%f:%f", a, b)]; ok {
				maxMap[fmt.Sprintf("%f:%f", a, b)]++
			} else {
				maxMap[fmt.Sprintf("%f:%f", a, b)] = 2
			}
			if max < maxMap[fmt.Sprintf("%f:%f", a, b)] {
				max = maxMap[fmt.Sprintf("%f:%f", a, b)]
			}
		}
	}
	return max
}

// @lc code=end

