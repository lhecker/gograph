// Copyright © 2014 The Gonum Authors, 2018 Leonard Hecker. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found at:
// https://github.com/gonum/gonum/blob/master/LICENSE

// The following file has been adapted to gograph's API in
// order to test the correctness of the Dijkstra algorithm.
package gograph

var ShortestPathTests = []struct {
	Name      string
	Arcs      []Arc
	Query     Arc
	NoPathFor Arc
	WantPaths [][]ID
}{
	{
		Name: "one edge directed",
		Arcs: []Arc{
			NewArc(0, 1, 1),
		},
		NoPathFor: NewArc(2, 3, 0),
		Query:     NewArc(0, 1, 0),
		WantPaths: [][]ID{
			{0, 1},
		},
	},
	{
		Name: "one edge self directed",
		Arcs: []Arc{
			NewArc(0, 1, 1),
		},
		NoPathFor: NewArc(2, 3, 0),
		Query:     NewArc(0, 0, 0),
	},
	{
		Name: "two paths directed",
		Arcs: []Arc{
			NewArc(0, 2, 2),
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
		},
		NoPathFor: NewArc(2, 1, 0),
		Query:     NewArc(0, 2, 0),
		WantPaths: [][]ID{
			{0, 1, 2},
			{0, 2},
		},
	},
	{
		Name: "confounding paths directed",
		Arcs: []Arc{
			// Add a path from 0->5 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 5, 1),
			// Add direct edge to goal of weight 4
			NewArc(0, 5, 4),
			// Add edge to a node that's still optimal
			NewArc(0, 2, 2),
			// Add edge to 3 that's overpriced
			NewArc(0, 3, 4),
			// Add very cheap edge to 4 which is a dead end
			NewArc(0, 4, 0),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 5, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 5},
			{0, 2, 3, 5},
			{0, 5},
		},
	},
	{
		Name: "confounding paths directed 2-step",
		Arcs: []Arc{
			// Add a path from 0->5 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 5, 1),
			// Add two step path to goal of weight 4
			NewArc(0, 6, 2),
			NewArc(6, 5, 2),
			// Add edge to a node that's still optimal
			NewArc(0, 2, 2),
			// Add edge to 3 that's overpriced
			NewArc(0, 3, 4),
			// Add very cheap edge to 4 which is a dead end
			NewArc(0, 4, 0),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 5, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 5},
			{0, 2, 3, 5},
			{0, 6, 5},
		},
	},
	{
		Name: "zero-weight cycle directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(1, 5, 0),
			NewArc(5, 1, 0),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
		},
	},
	{
		Name: "zero-weight cycle^2 directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(1, 5, 0),
			NewArc(5, 1, 0),
			// With its own zero-weight cycle.
			NewArc(5, 6, 0),
			NewArc(6, 5, 0),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
		},
	},
	{
		Name: "zero-weight cycle^2 confounding directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(1, 5, 0),
			NewArc(5, 1, 0),
			// With its own zero-weight cycle.
			NewArc(5, 6, 0),
			NewArc(6, 5, 0),
			// But leading to the target.
			NewArc(5, 4, 3),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
			{0, 1, 5, 4},
		},
	},
	{
		Name: "zero-weight cycle^3 directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(1, 5, 0),
			NewArc(5, 1, 0),
			// With its own zero-weight cycle.
			NewArc(5, 6, 0),
			NewArc(6, 5, 0),
			// With its own zero-weight cycle.
			NewArc(6, 7, 0),
			NewArc(7, 6, 0),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
		},
	},
	{
		Name: "zero-weight 3·cycle^2 confounding directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(1, 5, 0),
			NewArc(5, 1, 0),
			// With 3 of its own zero-weight cycles.
			NewArc(5, 6, 0),
			NewArc(6, 5, 0),
			NewArc(5, 7, 0),
			NewArc(7, 5, 0),
			// Each leading to the target.
			NewArc(5, 4, 3),
			NewArc(6, 4, 3),
			NewArc(7, 4, 3),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
			{0, 1, 5, 4},
			{0, 1, 5, 6, 4},
			{0, 1, 5, 7, 4},
		},
	},
	{
		Name: "zero-weight reversed 3·cycle^2 confounding directed",
		Arcs: []Arc{
			// Add a path from 0->4 of weight 4
			NewArc(0, 1, 1),
			NewArc(1, 2, 1),
			NewArc(2, 3, 1),
			NewArc(3, 4, 1),
			// Add a zero-weight cycle.
			NewArc(3, 5, 0),
			NewArc(5, 3, 0),
			// With 3 of its own zero-weight cycles.
			NewArc(5, 6, 0),
			NewArc(6, 5, 0),
			NewArc(5, 7, 0),
			NewArc(7, 5, 0),
			// Each leading from the source.
			NewArc(0, 5, 3),
			NewArc(0, 6, 3),
			NewArc(0, 7, 3),
		},
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
			{0, 5, 3, 4},
			{0, 6, 5, 3, 4},
			{0, 7, 5, 3, 4},
		},
	},
	{
		Name: "zero-weight |V|·cycle^(n/|V|) directed",
		Arcs: func() []Arc {
			e := []Arc{
				// Add a path from 0->4 of weight 4
				NewArc(0, 1, 1),
				NewArc(1, 2, 1),
				NewArc(2, 3, 1),
				NewArc(3, 4, 1),
			}

			next := len(e) + 1

			// Add n zero-weight cycles.
			for i := 0; i < 100; i++ {
				e = append(e,
					NewArc(next+i, i, 0),
					NewArc(i, next+i, 0),
				)
			}

			return e
		}(),
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
		},
	},
	{
		Name: "zero-weight n·cycle directed",
		Arcs: func() []Arc {
			e := []Arc{
				// Add a path from 0->4 of weight 4
				NewArc(0, 1, 1),
				NewArc(1, 2, 1),
				NewArc(2, 3, 1),
				NewArc(3, 4, 1),
			}

			next := len(e) + 1

			// Add n zero-weight cycles.
			for i := 0; i < 100; i++ {
				e = append(e,
					NewArc(next+i, 1, 0),
					NewArc(1, next+i, 0),
				)
			}

			return e
		}(),
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
		},
	},
	{
		Name: "zero-weight bi-directional tree with single exit directed",
		Arcs: func() []Arc {
			e := []Arc{
				// Add a path from 0->4 of weight 4
				NewArc(0, 1, 1),
				NewArc(1, 2, 1),
				NewArc(2, 3, 1),
				NewArc(3, 4, 1),
			}

			// Make a bi-directional tree rooted at node 2 with
			// a single exit to node 4 and co-equal cost from
			// 2 to 4.
			const (
				depth     = 4
				branching = 4
			)

			next := len(e) + 1
			src := 2
			last := 0

			for l := 0; l < depth; l++ {
				for i := 0; i < branching; i++ {
					last = next + i
					e = append(e,
						NewArc(src, last, 0),
						NewArc(last, src, 0),
					)
				}
				src = next + 1
				next += branching
			}

			e = append(e, NewArc(last, 4, 2))
			return e
		}(),
		NoPathFor: NewArc(4, 5, 0),
		Query:     NewArc(0, 4, 0),
		WantPaths: [][]ID{
			{0, 1, 2, 3, 4},
			{0, 1, 2, 6, 10, 14, 20, 4},
		},
	},
}
