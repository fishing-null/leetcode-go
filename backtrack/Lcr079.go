package backtrack

func subsets(nums []int) [][]int {
	ret = make([][]int, 0)
	track = make([]int, 0)
	subsetsBacktrack(nums, 0)
	return ret
}

func subsetsBacktrack(nums []int, start int) {
	tmp := make([]int, len(track))
	copy(tmp, track)
	ret = append(ret, tmp)
	for i := start; i < len(nums); i++ {
		track = append(track, nums[i])
		subsetsBacktrack(nums, i+1)
		track = track[:len(track)-1]
	}
}
