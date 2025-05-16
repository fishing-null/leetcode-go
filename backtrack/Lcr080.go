package backtrack

var (
	ret   [][]int
	track []int
)

func combine(n int, k int) [][]int {
	ret = make([][]int, 0)
	track = make([]int, 0)
	combineBacktrack(n, k, 1)
	return ret
}

func combineBacktrack(n int, k int, start int) {
	//只在叶子节点添加
	if k == len(track) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		ret = append(ret, tmp)
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		combineBacktrack(n, k, i+1)
		track = track[:len(track)-1]
	}
}
