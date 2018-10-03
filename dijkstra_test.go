package gograph

import (
	"reflect"
	"testing"
)

func TestNumericRosettacode(t *testing.T) {
	graph := numericRosettacodeGraph()
	path := Dijkstra(graph, 1, 5)

	if !reflect.DeepEqual(path, []ID{1, 3, 4, 5}) {
		t.Fatalf("invalid path: %v", path)
	}
}

func TestAlphabeticalRosettacode(t *testing.T) {
	graph := alphabeticalRosettacodeGraph()
	path := Dijkstra(graph, "a", "e")

	if !reflect.DeepEqual(path, []ID{"a", "c", "d", "e"}) {
		t.Fatalf("invalid path: %v", path)
	}
}

func TestGonumSuite(t *testing.T) {
	for _, test := range ShortestPathTests {
		test := test

		t.Run(test.Name, func(t *testing.T) {
			graph := NewDirectedGraph()

			for _, arc := range test.Arcs {
				graph.AddNode(NewNode(arc.GraphSourceID()))
				graph.AddNode(NewNode(arc.GraphTargetID()))
				graph.AddArc(arc)
			}

			path := Dijkstra(graph, test.NoPathFor.GraphSourceID(), test.NoPathFor.GraphTargetID())
			if len(path) != 0 {
				t.Errorf("expected no path, but found: %v", path)
				return
			}

			path = Dijkstra(graph, test.Query.GraphSourceID(), test.Query.GraphTargetID())
			ok := len(test.WantPaths) == 0 && len(path) == 0

			for _, expectedPath := range test.WantPaths {
				if reflect.DeepEqual(path, expectedPath) {
					ok = true
					break
				}
			}

			if !ok {
				t.Errorf("got %v, but expected either one of these: %v", path, test.WantPaths)
			}
		})
	}
}

func BenchmarkNumericRosettacode(b *testing.B) {
	graph := numericRosettacodeGraph()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Dijkstra(graph, 1, 5)
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
