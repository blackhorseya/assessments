package q3

// Solution resolve problem
func Solution(S string) int {
	ret, count, prev := 0, 0, int32(0)

	for fast, slow := 0, 0; fast < len(S); fast++ {
		ch := S[fast]
		if fast == 0 || prev != int32(ch) {
			prev = int32(ch)
			count = 1
		} else {
			count++
			if count >= 3 {
				slow = fast - 1
			}
		}

		ret = max(ret, fast-slow+1)
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
