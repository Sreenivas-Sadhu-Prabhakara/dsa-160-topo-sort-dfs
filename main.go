package main

import "fmt"

// TopoSort returns a topological order via DFS finish times: a vertex
// finishes only after everything it points to, so the reversed finish order
// is topological. ok is false when a cycle exists.
func TopoSort(n int, edges [][2]int) ([]int, bool) {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	const (
		white = 0
		gray  = 1
		black = 2
	)
	state := make([]int, n)
	out := make([]int, 0, n)
	ok := true
	var walk func(v int)
	walk = func(v int) {
		state[v] = gray
		for _, w := range adj[v] {
			switch state[w] {
			case gray:
				ok = false // back edge: cycle, no valid order
			case white:
				walk(w)
			}
		}
		state[v] = black
		out = append(out, v) // finished: everything below is already out
	}
	for s := 0; s < n && ok; s++ {
		if state[s] == white {
			walk(s)
		}
	}
	if !ok {
		return nil, false
	}
	for l, r := 0, len(out)-1; l < r; l, r = l+1, r-1 {
		out[l], out[r] = out[r], out[l]
	}
	return out, true
}

func main() {
	order, ok := TopoSort(6, [][2]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}})
	fmt.Println(order, ok)
}
