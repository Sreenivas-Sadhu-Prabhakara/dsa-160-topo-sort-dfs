package main

import "testing"

func validOrder(n int, edges [][2]int, order []int) bool {
	if len(order) != n {
		return false
	}
	pos := make([]int, n)
	for i, v := range order {
		pos[v] = i
	}
	for _, e := range edges {
		if pos[e[0]] >= pos[e[1]] {
			return false
		}
	}
	return true
}

func TestTopoSort(t *testing.T) {
	edges := [][2]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}
	order, ok := TopoSort(6, edges)
	if !ok || !validOrder(6, edges, order) {
		t.Errorf("order %v invalid", order)
	}

	if _, ok := TopoSort(3, [][2]int{{0, 1}, {1, 2}, {2, 0}}); ok {
		t.Error("cyclic graph must report no order")
	}
	if _, ok := TopoSort(1, [][2]int{{0, 0}}); ok {
		t.Error("self loop must report no order")
	}

	order, ok = TopoSort(4, [][2]int{{3, 2}, {2, 1}, {1, 0}})
	if !ok || !validOrder(4, [][2]int{{3, 2}, {2, 1}, {1, 0}}, order) {
		t.Errorf("chain order %v invalid", order)
	}

	order, ok = TopoSort(2, nil)
	if !ok || len(order) != 2 {
		t.Errorf("edgeless = %v", order)
	}
}
