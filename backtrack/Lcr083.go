package backtrack

var (
	ret   [][]int
	track []int
)

func permute(nums []int) [][]int {
	ret = make([][]int, 0)
	track = make([]int, 0)
	permuteBacktrack(nums, len(nums))
	return ret
}

func permuteBacktrack(nums []int, n int) {
	if len(track) == n {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ret = append(ret, tmp)
	}
	for i := range nums {
		if contains(track, nums[i]) {
			continue
		}
		track = append(track, nums[i])
		permuteBacktrack(nums, n)
		track = track[:len(track)-1]
	}
}

func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
