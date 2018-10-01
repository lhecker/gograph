package gograph

import (
	"container/heap"
)

const (
	maxInt = int(^uint(0) >> 1)
)

type dijkstraDistance struct {
	ID       ID
	Distance int
}

type dijkstraDistanceHeap []dijkstraDistance

func (h dijkstraDistanceHeap) Len() int           { return len(h) }
func (h dijkstraDistanceHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h dijkstraDistanceHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *dijkstraDistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(dijkstraDistance))
}

func (h *dijkstraDistanceHeap) Pop() interface{} {
	idx := len(*h) - 1
	tail := (*h)[idx]
	*h = (*h)[:idx]
	return tail
}

func (h *dijkstraDistanceHeap) updateDistance(id ID, val int) {
	for i := 0; i < len(*h); i++ {
		if (*h)[i].ID == id {
			(*h)[i].Distance = val
			heap.Fix(h, i)
			break
		}
	}
}

type dijkstraVertexData struct {
	Previous ID
	Distance int
	Visited  bool
}

func Dijkstra(graph *Graph, source ID, target ID) []ID {
	if _, ok := graph.vertices[source]; !ok {
		return nil
	}
	if _, ok := graph.vertices[target]; !ok {
		return nil
	}

	vertexData := map[ID]*dijkstraVertexData{
		source: {
			Distance: 0,
		},
	}

	Q := make(dijkstraDistanceHeap, 0, len(graph.vertices))
	Q = append(Q, dijkstraDistance{
		ID:       source,
		Distance: 0,
	})

	for id := range graph.vertices {
		if id == source {
			continue
		}

		vertexData[id] = &dijkstraVertexData{
			Distance: maxInt,
		}

		Q = append(Q, dijkstraDistance{
			ID:       id,
			Distance: maxInt,
		})
	}

	for Q.Len() != 0 {
		u := heap.Pop(&Q).(dijkstraDistance)
		if u.ID == target {
			break
		}

		udata := vertexData[u.ID]
		distance := udata.Distance

		udata.Visited = true

		for v, arc := range graph.arcs[u.ID] {
			vdata := vertexData[v]
			if vdata.Visited {
				continue
			}

			alt := distance + arc.Weight()

			if alt < vdata.Distance {
				vdata.Distance = alt
				vdata.Previous = u.ID
				Q.updateDistance(v, alt)
			}
		}
	}

	path := []ID(nil)
	id := target

	for {
		udata := vertexData[id]
		if udata.Previous == nil {
			break
		}

		path = append(path, id)
		id = udata.Previous
	}

	if path == nil {
		return nil
	}

	path = append(path, source)

	for left, right := 0, len(path)-1; left < right; left, right = left+1, right-1 {
		path[left], path[right] = path[right], path[left]
	}

	return path
}
