package gograph

import (
	"reflect"
	"testing"
)

func TestBasic(t *testing.T) {
	graph := NewGraph()

	graph.AddVertex(NewSimpleVertex(0))
	graph.AddVertex(NewSimpleVertex(1))
	graph.AddVertex(NewSimpleVertex(2))
	graph.AddVertex(NewSimpleVertex(3))
	graph.AddVertex(NewSimpleVertex(4))
	graph.AddVertex(NewSimpleVertex(5))

	graph.AddArc(NewSimpleArc(0, 1, 7))
	graph.AddArc(NewSimpleArc(0, 2, 9))
	graph.AddArc(NewSimpleArc(0, 5, 14))
	graph.AddArc(NewSimpleArc(1, 2, 10))
	graph.AddArc(NewSimpleArc(1, 3, 15))
	graph.AddArc(NewSimpleArc(2, 3, 11))
	graph.AddArc(NewSimpleArc(2, 5, 2))
	graph.AddArc(NewSimpleArc(3, 4, 6))
	graph.AddArc(NewSimpleArc(4, 5, 9))

	path, _ := Dijkstra(graph, 0, 4)

	if !reflect.DeepEqual(path, []ID{0, 2, 3, 4}) {
		t.Fatalf("invalid path: %v", path)
	}
}
