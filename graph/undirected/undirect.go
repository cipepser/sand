package undirected

// Graph is a graph with n vertices and edges.
type Graph struct {
	n     int
	edges [][]int
}

// NewGraph creates a new graph with n vertices.
func NewGraph(n int) *Graph {
	g := &Graph{
		n:     n,
		edges: make([][]int, n),
	}
	return g
}

// AddEdge adds a edge connects vertex u to v and v to u.
func (g *Graph) AddEdge(u, v int) {
	g.edges[v] = append(g.edges[v], u)
	g.edges[u] = append(g.edges[u], v)
}
