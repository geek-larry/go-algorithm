package _5__颜色分类

// 双指针
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

// 双指针
func sortColors2(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for ; i <= p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
