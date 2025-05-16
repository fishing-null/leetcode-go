package binarysearch

func lengthOfLIS(nums []int) int {
	arr := make([]int, 0)
	for _, v := range nums {
		left, right := 0, len(arr)
		for left < right {
			mid := left + (right-left)/2
			if arr[mid] < v {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == len(arr) {
			arr = append(arr, v)
		} else {
			arr[left] = v
		}
	}
	return len(arr)
}
