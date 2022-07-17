package day58_152_maxProduct

func maxProduct(nums []int) int {
	if len(nums) < 1 {
		return 0
	}
	dpMax, dpMin := make([]int, len(nums)), make([]int, len(nums))
	dpMax[0], dpMin[0] = nums[0], nums[0]
	res := dpMax[0]
	for i := 1; i < len(nums); i++ {
		dpMax[i] = max(dpMax[i-1]*nums[i], max(dpMin[i-1]*nums[i], nums[i]))
		if dpMax[i] > res {
			res = dpMax[i]
		}
		dpMin[i] = min(dpMin[i-1]*nums[i], min(dpMax[i-1]*nums[i], nums[i]))
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
