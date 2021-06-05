package q2

import (
	"strconv"
)

// Solution resolve problem
func Solution(S string) int {
	ret := 0

	for i := 0; i < len(S)-1; i++ {
		num, _ := strconv.Atoi(string(S[i]) + string(S[i+1]))

		if num > ret {
			ret = num
		}
	}

	return ret
}
