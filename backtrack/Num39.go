package backtrack

func combinationSum(candidates []int, target int) [][]int {
	ret = make([][]int, 0)
	track = make([]int, 0)
	backtrackCombinationSum(target, 0, candidates)
	return ret
}

func backtrackCombinationSum(target, start int, candidates []int) {
	//组合数==target时 添加到ret中
	sum := 0
	for i := range track {
		sum += track[i]
	}
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ret = append(ret, tmp)
	}
	if sum > target {
		return
	}
	for i := start; i < len(candidates); i++ {
		track = append(track, candidates[i])
		backtrackCombinationSum(target, i, candidates)
		track = track[:len(track)-1]
	}
}
