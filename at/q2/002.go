package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func q2(nums []int, target int) []int {
	left, right, ret := 0, len(nums)-1, -1

	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			ret = max(ret, sum)
			left++
		} else {
			right--
		}
	}

	if ret == -1 {
		return nil
	}

	return []int{left, right + 1}
}
