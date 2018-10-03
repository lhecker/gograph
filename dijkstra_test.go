package gograph

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	graph := NewDirectedGraph()

	graph.AddNode(NewNode(0))
	graph.AddNode(NewNode(1))
	graph.AddNode(NewNode(2))
	graph.AddNode(NewNode(3))
	graph.AddNode(NewNode(4))
	graph.AddNode(NewNode(5))

	graph.AddArc(NewArc(0, 1, 7))
	graph.AddArc(NewArc(0, 2, 9))
	graph.AddArc(NewArc(0, 5, 14))
	graph.AddArc(NewArc(1, 2, 10))
	graph.AddArc(NewArc(1, 3, 15))
	graph.AddArc(NewArc(2, 3, 11))
	graph.AddArc(NewArc(2, 5, 2))
	graph.AddArc(NewArc(3, 4, 6))
	graph.AddArc(NewArc(4, 5, 9))

	path := Dijkstra(graph, 0, 4)

	if !reflect.DeepEqual(path, []ID{0, 2, 3, 4}) {
		t.Fatalf("invalid path: %v", path)
	}
}
