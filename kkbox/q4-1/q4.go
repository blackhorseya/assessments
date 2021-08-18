package q4_1

func Solution(A []int) int {
	if len(A) >= 100000 {
		return -1
	}

	ret := 0
	for i := 0; i < len(A); i++ {
		for j := i; j < len(A); j++ {
			sum := 0

			for k := i; k <= j; k++ {
				sum += A[k]
			}

			if sum == 0 {
				ret++
			}
		}
	}

	return ret
}
