package main

import "testing"

type Edge struct {
	from string
	to   string
}

func initGraph() {
	graph = make(map[string]map[string]bool)
}

func TestAddEdgeTestAddEdge(t *testing.T) {
	tests := []struct {
		input []Edge
		want  map[string]map[string]bool
	}{
		{
			input: []Edge{
				{
					from: "this",
					to:   "is",
				},
				{
					from: "this",
					to:   "was",
				},
				{
					from: "they",
					to:   "are",
				},
			},
			want: map[string]map[string]bool{
				"this": {
					"is":  true,
					"was": true,
				},
				"they": {
					"are": true,
				},
			},
		},
	}

	for _, tt := range tests {
		for _, edge := range tt.input {
			addEdge(edge.from, edge.to)
		}

		if len(graph) != len(tt.want) {
			t.Errorf("len(graph) = %d, want = %d", len(graph), len(tt.want))
		}
		for from, to := range graph {
			if tt.want[from] == nil {
				t.Errorf("tt.want[%s] = nil, want not nil", from)
			}
			if len(to) != len(tt.want[from]) {
				t.Errorf("len(graph[%s]) = %d, want = %d", from, len(to), len(tt.want[from]))
			}
			for k, v := range to {
				if !v || !tt.want[from][k] {
					t.Errorf("graph[from][k] = %v, want = true", v)
				}
			}
		}

		initGraph()
	}
}

func TestHasEdgeTestHasEdge(t *testing.T) {
	graph = map[string]map[string]bool{
		"this": {
			"is":  true,
			"was": true,
		},
		"they": {
			"are": true,
		},
	}
	tests := []struct {
		input Edge
		want  bool
	}{
		{
			input: Edge{
				from: "this",
				to:   "is",
			},
			want: true,
		},
		{
			input: Edge{
				from: "they",
				to:   "are",
			},
			want: true,
		},
		{
			input: Edge{
				from: "they",
				to:   "were",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		if hasEdge(tt.input.from, tt.input.to) != tt.want {
			t.Errorf(
				"hasEdge(%s, %s) = %v, want = %v",
				tt.input.from,
				tt.input.to,
				hasEdge(tt.input.from, tt.input.to),
				tt.want,
			)
		}
	}
}
