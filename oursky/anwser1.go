package oursky

func contains(arr []string, target string) bool {
	for _, ch := range arr {
		if ch == target {
			return true
		}
	}

	return false
}

func isSubset(set1 []string, set2 []string) bool {
	if len(set2) > len(set1) {
		return isSubset(set2, set1)
	}

	for _, ch2 := range set2 {
		if !contains(set1, ch2) {
			return false
		}
	}

	return true
}
