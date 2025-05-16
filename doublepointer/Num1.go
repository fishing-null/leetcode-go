package doublepointer

import "sort"

func twoSum(nums []int, target int) []int {
	// 存储值和原始索引
	type NumWithIndex struct {
		value int
		index int
	}
	numsWithIndex := make([]NumWithIndex, len(nums))
	for i, num := range nums {
		numsWithIndex[i] = NumWithIndex{num, i}
	}

	// 按值排序
	sort.Slice(numsWithIndex, func(i, j int) bool {
		return numsWithIndex[i].value < numsWithIndex[j].value
	})

	left, right := 0, len(nums)-1
	for left < right {
		sum := numsWithIndex[left].value + numsWithIndex[right].value
		if sum == target {
			// 返回原始索引（注意顺序，可能 left > right）
			return []int{numsWithIndex[left].index, numsWithIndex[right].index}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}
