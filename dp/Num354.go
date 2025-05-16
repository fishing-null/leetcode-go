package dp

import "sort"

func MaxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			//按照h降序
			return envelopes[i][1] > envelopes[j][1]
		} else {
			//按照w升序
			return envelopes[i][0] < envelopes[j][0]
		}
	})
	length := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		length[i] = envelopes[i][1]
	}
	return lengthOfLIS(length)
}
