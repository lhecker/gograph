package gograph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDijkstraUnreachableError(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))

	path, distance, err := Dijkstra(graph, 1, 2)
	assert.EqualError(err, "no path found for 1 -> 2")
	assert.Equal(infinity, distance)
	assert.Empty(path)
}

func TestDijkstraUnknownNodeErrorForSource(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))

	path, distance, err := Dijkstra(graph, 9, 5)
	assert.EqualError(err, "unknown node 9")
	assert.Equal(infinity, distance)
	assert.Empty(path)
}

func TestDijkstraUnknownNodeErrorForTarget(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))

	path, distance, err := Dijkstra(graph, 1, 9)
	assert.EqualError(err, "unknown node 9")
	assert.Equal(infinity, distance)
	assert.Empty(path)
}

func TestDijkstraUnknownNodeErrorForArcs(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(3))
	graph.AddArc(NewArc(1, 2, 1))
	graph.AddArc(NewArc(2, 3, 1))

	path, distance, err := Dijkstra(graph, 1, 3)
	assert.EqualError(err, "unknown node 2")
	assert.Equal(infinity, distance)
	assert.Empty(path)
}

func TestDijkstraNegativeWeightError(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))
	graph.AddArc(NewArc(1, 2, -1))

	path, distance, err := Dijkstra(graph, 1, 2)
	assert.EqualError(err, "negative weight -1 on arc 1 -> 2")
	assert.Equal(infinity, distance)
	assert.Empty(path)
}

func TestDijkstraWhenSourceIsTarget(t *testing.T) {
	assert := assert.New(t)

	graph := NewDirectedGraph()
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))
	graph.AddArc(NewArc(1, 1, 1))
	graph.AddArc(NewArc(1, 2, 1))
	graph.AddArc(NewArc(2, 1, 1))

	path, distance, err := Dijkstra(graph, 1, 1)
	assert.NoError(err)
	assert.Zero(distance)
	assert.Empty(path)
}

func TestDijkstraWithNumericRosettacode(t *testing.T) {
	assert := assert.New(t)
	graph := numericRosettacodeGraph()

	path, distance, err := Dijkstra(graph, 1, 5)
	if assert.NoError(err) {
		assert.Equal(26.0, distance)
		assert.EqualValues(path, []ID{1, 3, 4, 5})
	}
}

func TestDijkstraWithAlphabeticalRosettacode(t *testing.T) {
	assert := assert.New(t)
	graph := alphabeticalRosettacodeGraph()

	path, distance, err := Dijkstra(graph, "a", "e")
	if assert.NoError(err) {
		assert.Equal(26.0, distance)
		assert.EqualValues(path, []ID{"a", "c", "d", "e"})
	}
}

func TestDijkstraWithGonumSuite(t *testing.T) {
	for _, test := range ShortestPathTests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			assert := assert.New(t)
			graph := NewDirectedGraph()

			for _, arc := range test.Arcs {
				graph.AddNode(NewNode(arc.GraphSourceID()))
				graph.AddNode(NewNode(arc.GraphTargetID()))
				graph.AddArc(arc)
			}

			path, distance, err := Dijkstra(graph, test.NoPathFor.GraphSourceID(), test.NoPathFor.GraphTargetID())
			assert.Error(err)
			assert.Equal(infinity, distance)
			assert.Empty(path)

			path, _, err = Dijkstra(graph, test.Query.GraphSourceID(), test.Query.GraphTargetID())
			if assert.NoError(err) {
				if len(test.WantPaths) == 0 {
					assert.Empty(path)
				} else {
					assert.Contains(test.WantPaths, path)
				}
			}
		})
	}
}

func BenchmarkNumericRosettacode(b *testing.B) {
	graph := numericRosettacodeGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _, _ = Dijkstra(graph, 1, 5)
	}
}

func numericRosettacodeGraph() *BuiltinGraph {
	graph := NewDirectedGraph()

	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))
	graph.AddNode(NewNode(3))
	graph.AddNode(NewNode(4))
	graph.AddNode(NewNode(5))
	graph.AddNode(NewNode(6))

	graph.AddArc(NewArc(1, 2, 7))
	graph.AddArc(NewArc(1, 3, 9))
	graph.AddArc(NewArc(1, 6, 14))
	graph.AddArc(NewArc(2, 3, 10))
	graph.AddArc(NewArc(2, 4, 15))
	graph.AddArc(NewArc(3, 4, 11))
	graph.AddArc(NewArc(3, 6, 2))
	graph.AddArc(NewArc(4, 5, 6))
	graph.AddArc(NewArc(5, 6, 9))

	return graph
}

func alphabeticalRosettacodeGraph() *BuiltinGraph {
	graph := NewDirectedGraph()

	graph.AddNode(NewNode("a"))
	graph.AddNode(NewNode("b"))
	graph.AddNode(NewNode("c"))
	graph.AddNode(NewNode("d"))
	graph.AddNode(NewNode("e"))
	graph.AddNode(NewNode("f"))

	graph.AddArc(NewArc("a", "b", 7))
	graph.AddArc(NewArc("a", "c", 9))
	graph.AddArc(NewArc("a", "f", 14))
	graph.AddArc(NewArc("b", "c", 10))
	graph.AddArc(NewArc("b", "d", 15))
	graph.AddArc(NewArc("c", "d", 11))
	graph.AddArc(NewArc("c", "f", 2))
	graph.AddArc(NewArc("d", "e", 6))
	graph.AddArc(NewArc("e", "f", 9))

	return graph
}
