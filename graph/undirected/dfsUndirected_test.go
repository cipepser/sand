package undirected

import (
	"testing"
)

func TestGraph_ExistsCycle(t *testing.T) {
	type fields struct {
		n     int
		edges [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "exists cycle",
			fields: fields{
				n: 4,
				edges: [][]int{
					[]int{1, 2, 3},
					[]int{0, 3},
					[]int{0},
					[]int{0, 1},
				},
			},
			want: true,
		},
		{
			name: "does not exist cycle",
			fields: fields{
				n: 4,
				edges: [][]int{
					[]int{1, 2, 3},
					[]int{0},
					[]int{0},
					[]int{0},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				n:     tt.fields.n,
				edges: tt.fields.edges,
			}
			if got := g.ExistsCycle(); got != tt.want {
				t.Errorf("Graph.ExistsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dfs(t *testing.T) {
	type args struct {
		p       int
		c       int
		edges   [][]int
		visited map[int]struct{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only a root vertex",
			args: args{
				p:       -1,
				c:       0,
				edges:   make([][]int, 1),
				visited: make(map[int]struct{}),
			},
			want: false,
		},
		{
			name: "exists a cycle(already visited)",
			args: args{
				p:     0,
				c:     0,
				edges: make([][]int, 0),
				visited: map[int]struct{}{
					0: struct{}{},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.p, tt.args.c, tt.args.edges, tt.args.visited); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
