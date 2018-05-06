package undirected

func (g *Graph) ExistsCycle() bool {
	return dfs(-1, 0, g.edges, map[int]struct{}{})
}

func dfs(p, c int, edges [][]int, visited map[int]struct{}) bool {
	_, ok := visited[c]
	if ok {
		return true
	}
	visited[c] = struct{}{}

	for _, v := range edges[c] {
		if v == p {
			continue
		}
		return dfs(c, v, edges, visited)
	}
	return false
}
