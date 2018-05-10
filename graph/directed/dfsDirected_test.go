package directed

import "testing"

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
				n: 6,
				edges: [][]int{
					[]int{1, 2},
					[]int{3, 4},
					[]int{3, 4},
					[]int{5},
					[]int{},
					[]int{2},
				},
			},
			want: true,
		},
		{
			name: "dose not exist cycle",
			fields: fields{
				n: 6,
				edges: [][]int{
					[]int{1, 2},
					[]int{3, 4},
					[]int{4},
					[]int{5},
					[]int{},
					[]int{2},
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
		c     int
		wSet  map[int]struct{}
		gSet  map[int]struct{}
		bSet  map[int]struct{}
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "only a root vertex",
			args: args{
				c: 0,
				wSet: map[int]struct{}{
					0: struct{}{},
				},
				gSet:  make(map[int]struct{}),
				bSet:  make(map[int]struct{}),
				edges: make([][]int, 1),
			},
			want: false,
		},
		{
			name: "exists a cycle(already visited)",
			args: args{
				c: 1,
				wSet: map[int]struct{}{
					1: struct{}{},
				},
				gSet: map[int]struct{}{
					0: struct{}{},
				},
				bSet: make(map[int]struct{}),
				edges: [][]int{
					[]int{},
					[]int{0},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dfs(tt.args.c, tt.args.wSet, tt.args.gSet, tt.args.bSet, tt.args.edges); got != tt.want {
				t.Errorf("dfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
