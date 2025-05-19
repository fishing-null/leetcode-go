package backtrack

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret = make([][]int, 0)
	track = make([]int, 0)
	backtrackSubsetsWithDup(nums, 0)
	return ret
}

func backtrackSubsetsWithDup(nums []int, start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	ret = append(ret, tmp)
	for i := start; i < len(nums); i++ {
		//判断是否重复
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		track = append(track, nums[i])
		backtrackSubsetsWithDup(nums, i+1)
		track = track[:len(track)-1]
	}
	return
}
