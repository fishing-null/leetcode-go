package dp

func lengthOfLIS(nums []int) int {
	//dp数组 dp[i]表示以nums[i]结尾所能得到的最长子序列
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	ret := 0
	for i := 0; i < len(dp); i++ {
		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}
