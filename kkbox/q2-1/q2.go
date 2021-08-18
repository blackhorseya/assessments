package q2_1

func Solution(A []int) int {
	sum := 1
	for _, val := range A {
		sum = sum * val
	}

	if sum > 0 {
		return 1
	} else if sum < 0 {
		return -1
	} else {
		return 0
	}
}
