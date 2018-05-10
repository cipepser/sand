package directed

func (g *Graph) ExistsCycle() bool {
	wSet := make(map[int]struct{}, g.n)
	gSet := make(map[int]struct{}, g.n)
	bSet := make(map[int]struct{}, g.n)

	for i := 0; i < g.n; i++ {
		wSet[i] = struct{}{}
	}

	for len(wSet) != 0 {
		c := -1
		for c = range wSet {
			break
		}
		if dfs(c, wSet, gSet, bSet, g.edges) {
			return true
		}
	}
	return false
}

func dfs(c int, wSet, gSet, bSet map[int]struct{}, edges [][]int) bool {
	delete(wSet, c)
	gSet[c] = struct{}{}

	for _, n := range edges[c] {
		_, ok := bSet[n]
		if ok {
			continue
		}
		_, ok = gSet[n]
		if ok {
			return true
		}
		if dfs(n, wSet, gSet, bSet, edges) {
			return true
		}
	}

	delete(gSet, c)
	bSet[c] = struct{}{}
	return false
}
