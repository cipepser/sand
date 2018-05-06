package undirected

import (
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *Graph
	}{
		{
			name: "create a new graph with 3 vertices",
			args: args{3},
			want: &Graph{
				n:     3,
				edges: make([][]int, 3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGraph(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGraph_AddEdge(t *testing.T) {
	type fields struct {
		n     int
		edges [][]int
	}
	type args struct {
		u int
		v int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Graph{
				n:     tt.fields.n,
				edges: tt.fields.edges,
			}
			g.AddEdge(tt.args.u, tt.args.v)
		})
	}
}
