/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	s      []int
	length int
	min    int
}

func Constructor() MinStack {
	return MinStack{
		s:   make([]int, 0),
		min: int(^uint(0) >> 1),
	}
}

func (this *MinStack) Push(val int) {
	this.s = append(this.s, val)
	this.length++
	if val < this.min {
		this.min = val
	}
}

func (this *MinStack) Pop() {
	this.length--
	if this.s[this.length] == this.min {
		this.min = int(^uint(0) >> 1)
		for _, tmp := range this.s[:this.length] {
			if tmp < this.min {
				this.min = tmp
			}
		}
	}
	this.s = this.s[:this.length]
}

func (this *MinStack) Top() int {
	return this.s[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

