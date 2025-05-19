package arr

import "sort"

func triangleType(nums []int) string {
	//sort
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	if nums[2]+nums[1] > nums[0] {
		if nums[0] == nums[1] && nums[1] == nums[2] {
			return "equilateral"
		}
		if nums[0] == nums[1] || nums[1] == nums[2] {
			return "isosceles "
		}
		return "scalene"
	} else {
		return "none"
	}
}
