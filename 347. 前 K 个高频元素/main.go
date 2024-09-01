package _47__前_K_个高频元素

import (
	"math/rand"
	"time"
)

func topKFrequent(nums []int, k int) []int {
	occurrences := map[int]int{}
	// 获取每个数字出现次数
	for _, num := range nums {
		occurrences[num]++
	}
	values := [][]int{}
	for key, value := range occurrences {
		values = append(values, []int{key, value})
	}
	ret := make([]int, k)
	qsort(values, 0, len(values)-1, ret, 0, k)
	return ret
}

func qsort(values [][]int, start, end int, ret []int, retIndex, k int) {
	rand.Seed(time.Now().UnixNano())
	picked := rand.Int()%(end-start+1) + start
	values[picked], values[start] = values[start], values[picked]

	pivot := values[start][1]
	index := start

	for i := start + 1; i <= end; i++ {
		// 使用双指针把不小于基准值的元素放到左边，
		// 小于基准值的元素放到右边
		if values[i][1] >= pivot {
			values[index+1], values[i] = values[i], values[index+1]
			index++
		}
	}
	values[start], values[index] = values[index], values[start]
	if k <= index-start {
		// 前 k 大的值在左侧的子数组里
		qsort(values, start, index-1, ret, retIndex, k)
	} else {
		// 前 k 大的值等于左侧的子数组全部元素
		// 加上右侧子数组中前 k - (index - start + 1) 大的值
		for i := start; i <= index; i++ {
			ret[retIndex] = values[i][0]
			retIndex++
		}
		if k > index-start+1 {
			qsort(values, index+1, end, ret, retIndex, k-(index-start+1))
		}
	}
}
