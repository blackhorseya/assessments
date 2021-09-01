package q3

// Solution ...
func Solution(riddle string) string {
	b := []byte(riddle)
	var prev byte
	for i, ch := range b {
		if b[i] == '?' {
			var next byte
			if i != len(riddle)-1 {
				next = b[i+1]
			}

			if prev != 'a' && next != 'a' {
				b[i] = 'a'
			} else if prev != 'b' && next != 'b' {
				b[i] = 'b'
			} else {
				b[i] = 'c'
			}
		}

		prev = ch
	}

	return string(b)
}
