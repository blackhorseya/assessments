package oursky

func recur(n float64, cur float64) float64 {
	if n < 2 {
		panic("invalid input")
	}

	if n == 2 {
		return 1/n + cur
	}

	return recur(n-1, cur+1/(n*(n-1)))
}
