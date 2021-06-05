package q4

// Tree declare a tree struct
type Tree struct {
	X int
	L *Tree
	R *Tree
}

func traversal(t *Tree, m map[int]int) int {
	if t == nil ||  m[t.X] != 0{
		return 0
	}

	m[t.X]++

	length := max(traversal(t.L, m) + 1, traversal(t.R, m) + 1)

	m[t.X]--

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Solution resolve problem
func Solution(T *Tree) int {
	return traversal(T, make(map[int]int))
}
